package verifier

import "net"

type Mx struct {
	Records     []*net.MX
	HasMxRecord bool
}

func isLenRecords(records []*net.MX) bool {
	return len(records) > 0
}

func (v *Verifier) CheckMx(domain string) (*Mx, error) {
	domain = domainToASCII(domain)
	mx, err := net.LookupMX(domain)
	if err != nil && !isLenRecords(mx) {
		return nil, err
	}

	return &Mx{
		Records:     mx,
		HasMxRecord: isLenRecords(mx),
	}, nil
}
