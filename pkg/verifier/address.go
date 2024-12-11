package verifier

import (
	"errors"
	"regexp"
	"strings"
)

const emailRegexString = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

var emailRegex = regexp.MustCompile(emailRegexString)

type Syntax struct {
	Username string
	Domain   string
	Valid    bool
}

type Result struct {
	Syntax Syntax
	Email  string
}

func isAddressValid(email string) bool {
	return emailRegex.MatchString(email)
}

func ParseAddress(email string) *Syntax {

	isAddressValid := isAddressValid(email)
	if !isAddressValid {
		return &Syntax{Valid: false}
	}

	index := strings.LastIndex(email, "@")
	username := email[:index]
	domain := email[index+1:]

	return &Syntax{
		Username: username,
		Domain:   domain,
		Valid:    isAddressValid,
	}
}

func VerifyEmail(email string) (*Result, error) {
	syntax := ParseAddress(email)
	res := Result{
		Email: email,
	}

	res.Syntax = *syntax

	if !syntax.Valid {
		return &res, errors.New("некорректный email")
	}

	return &res, nil
}
