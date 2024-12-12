package verifier

import (
	"net/smtp"
)

type Verifier struct {
	Syntax *Syntax
	Mx     *Mx
	SMTP   *SMTP
}

func NewVierifier() *Verifier {
	return &Verifier{}
}

func (v *Verifier) Verify(email string) (*smtp.Client, error) {
	if err := v.ParseAddress(email); err != nil {
		return nil, err
	}

	if err := v.CheckMx(); err != nil {
		return nil, err
	}

	if err := v.CheckSMTP(); err != nil {
		return nil, err
	}

	return v.SMTP.Client, nil
}
