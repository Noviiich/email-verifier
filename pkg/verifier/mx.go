package verifier

import "net"

func CheckMx(domain string) (bool, error) {
	mxRecords, err := net.LookupMX(domain)
	if err != nil {
		return false, err
	}

	return len(mxRecords) > 0, nil
}
