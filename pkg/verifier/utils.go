package verifier

import "golang.org/x/net/idna"

func domainToASCII(domain string) string {
	asciiDomain, err := idna.ToASCII(domain)
	if err != nil {
		return domain
	}

	return asciiDomain
}
