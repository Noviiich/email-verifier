package verifier

import (
	"errors"
	"regexp"
	"strings"
)

var emailRegex = regexp.MustCompile(emailRegexString)

type Syntax struct {
	Email    string
	Username string
	Domain   string
	Valid    bool
}

func isAddressValid(email string) bool {
	return emailRegex.MatchString(email)
}

func (v *Verifier) ParseAddress(email string) error {

	isAddressValid := isAddressValid(email)
	if !isAddressValid {
		v.Syntax = &Syntax{
			Email: email,
			Valid: isAddressValid,
		}
		return errors.New("Введен некорректный email")
	}

	index := strings.LastIndex(email, "@")
	username := email[:index]
	domain := email[index+1:]

	v.Syntax = &Syntax{
		Email:    email,
		Username: username,
		Domain:   domain,
		Valid:    isAddressValid,
	}
	return nil
}
