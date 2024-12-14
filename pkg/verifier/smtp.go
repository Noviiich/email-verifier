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
	HostExists  bool //сузествует ли хост
	FullInbox   bool //заполнен ли электронный ящик у почты
	Deliverable bool //можно ли отправить электронно письмо на сервер
	CatchAll    bool //есть ли у домена универсальный адресс почты
	Disabled    bool //электронная почта заблокирована или отличена провайдером
}

func newSMTPClient(domain string) (*smtp.Client, error) {
	domain = domainToASCII(domain)
	mxRecords, err := net.LookupMX(domain)
	if err != nil {
		return nil, err
	}

	if len(mxRecords) == 0 {
		return nil, errors.New("no MX records found")
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
				return nil, errors.New("all MX records failed")
			}
		default:
			return nil, errors.New("unknown error")
		}

	}
}

func (v *Verifier) CheckSMTP(domain, username string) (*SMTP, error) {

	if !v.smtpCheckEnabled {
		return nil, nil
	}

	var smtp SMTP
	//email := fmt.Sprintf("%s@%s", username, domain)

	client, err := newSMTPClient(domain)
	if err != nil {
		return &smtp, err
	}

	defer client.Close()

	if err := client.Hello(defaultHelloName); err != nil {
		return &smtp, nil
	}

	if err := client.Mail(defaultFromEmail); err != nil {
		return &smtp, nil
	}

	smtp.HostExists = true
	smtp.CatchAll = true

	return &smtp, nil
}
