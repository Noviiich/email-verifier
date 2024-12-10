package verifier

import (
	"errors"
	"regexp"
)

const emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

func VerifyEmail(email string) (bool, error) {
	if email == "" {
		return false, errors.New("email is empty")
	}

	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email), nil
}
