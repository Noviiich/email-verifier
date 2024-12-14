package verifier

// служит для проверки почты
type Verifier struct {
	smtpCheckEnabled     bool   //Включена ли проверка SMTP(включена по умолчанию)
	catchAllCheckEnabled bool   //catch all - адрес, на который приходят письма от несуществующих пользователей
	fromEmail            string //Почта, использует в MAIL команде, "user@example.org" по умолчанию
	helloName            string //Имя, которое использует в Hello команде, "localhost" по умолчанию
}

type Result struct {
	Email      string
	Reachable  string //некоторая информация, которая показывает, является ли адрес электронной почты достижимым
	Disposable bool   //является ли одноразовым адресом почты
	Free       bool   // является ли domain бесплатным доменом электронной почты
	Syntax     Syntax //все подробности об адресе электронной почты
	Mx         *Mx
	SMTP       *SMTP
}

func NewVierifier() *Verifier {
	return &Verifier{
		smtpCheckEnabled:     true,
		catchAllCheckEnabled: true,
		fromEmail:            defaultFromEmail,
		helloName:            defaultHelloName,
	}
}

func (v *Verifier) Verify(email string) (*Result, error) {

	res := Result{
		Email: email,
	}

	syntax := v.ParseAddress(email)
	if !syntax.Valid {
		return &res, nil
	}
	res.Syntax = syntax
	domain := syntax.Domain

	mx, err := v.CheckMx(domain)
	if err != nil {
		return nil, err
	}
	res.Mx = mx

	smtp, err := v.CheckSMTP(domain, syntax.Username)
	if err != nil {
		return nil, err
	}
	res.SMTP = smtp

	return &res, nil
}
