package verifier

import (
	"regexp"
	"strings"
)

var emailRegex = regexp.MustCompile(emailRegexString)

type Syntax struct {
	Username string
	Domain   string
	Valid    bool
}

func isAddressValid(email string) bool {
	return emailRegex.MatchString(email)
}

func (v *Verifier) ParseAddress(email string) Syntax {

	isAddressValid := isAddressValid(email)
	if !isAddressValid {
		return Syntax{
			Valid: false,
		}
	}

	index := strings.LastIndex(email, "@")
	username := email[:index]
	domain := email[index+1:]

	return Syntax{
		Username: username,
		Domain:   domain,
		Valid:    isAddressValid,
	}
}
