package verifier

import (
	"errors"
	"regexp"
)

func isAddressValid(email string) bool {
	return emailRegex.MatchString(email)
}

const emailRegexString = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

var emailRegex = regexp.MustCompile(emailRegexString)

func VerifyEmail(email string) (bool, error) {
	if email == "" {
		return false, errors.New("email is empty")
	}

	isAddressValid := isAddressValid(email)
	return isAddressValid, nil
}
