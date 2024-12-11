package verifier

import "net"

type Mx struct {
	Records     []*net.MX
	HasMxRecord bool
}

func CheckMx(domain string) (*Mx, error) {
	domain = domainToASCII(domain)
	mxRecords, err := net.LookupMX(domain)
	if err != nil {
		return &Mx{HasMxRecord: false}, err
	}

	return &Mx{
		Records:     mxRecords,
		HasMxRecord: len(mxRecords) > 0,
	}, nil
}
