package verifier

import (
	"errors"
	"net"
	"net/smtp"
	"sync"
)

var (
	mutex sync.Mutex
	done  bool
)

type SMTP struct {
	HostExists bool
	Client     *smtp.Client
}

func newSMTPClient(domain string) (*smtp.Client, error) {
	domain = domainToASCII(domain)
	mxRecords, err := net.LookupMX(domain)
	if err != nil {
		return nil, err
	}

	if len(mxRecords) == 0 {
		return nil, errors.New("No MX records found")
	}

	ch := make(chan interface{}, 1)

	for _, r := range mxRecords {

		addr := r.Host + smtpPort

		go func() {
			c, err := smtp.Dial(addr)
			if err != nil {
				if !done {
					ch <- err
				}
				return
			}

			mutex.Lock()
			switch {
			case !done:
				done = true
				ch <- c
			default:
				c.Close()
			}
			mutex.Unlock()

		}()
	}

	var errs []error

	for {
		res := <-ch
		switch r := res.(type) {
		case *smtp.Client:
			return r, nil
		case error:
			errs = append(errs, r)
			if len(errs) == len(mxRecords) {
				return nil, errors.New("All MX records failed")
			}
		default:
			return nil, errors.New("Unknown error")
		}

	}
}

func (v *Verifier) CheckSMTP() error {
	client, err := newSMTPClient(v.Syntax.Domain)

	if err != nil {
		v.SMTP = &SMTP{HostExists: false}
		return err
	}

	v.SMTP = &SMTP{HostExists: true, Client: client}
	return nil
}
