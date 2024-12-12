package verifier

import "net"

type Mx struct {
	Records     []*net.MX
	HasMxRecord bool
}

func (v *Verifier) CheckMx() error {
	domain := domainToASCII(v.Syntax.Domain)
	mxRecords, err := net.LookupMX(domain)
	if err != nil {
		v.Mx = &Mx{
			HasMxRecord: false,
		}
		return err
	}

	v.Mx = &Mx{
		Records:     mxRecords,
		HasMxRecord: true,
	}
	return nil
}
