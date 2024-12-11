package verifier

import (
	"fmt"
	"net/smtp"
)

func CheckSMTP(domain, username string) (*smtp.Client, error) {
	mx, err := CheckMx(domain)
	if err != nil {
		fmt.Println("Нет такой mx-записи")
		return nil, err
	}
	var conn *smtp.Client

	for _, record := range mx.Records {
		addr := record.Host + smtpPort
		conn, err = smtp.Dial(addr)
		if err != nil {
			fmt.Printf("На сервере %s проблемы", record.Host)
			continue
		}

		if err := conn.Hello(defaultHelloName); err != nil {
			fmt.Printf("На сервере %s проблемы с отправкой hello", record.Host)
			continue
		}

		if err := conn.Mail(defaultFromEmail); err != nil {
			fmt.Printf("На сервере %s проблемы с отправкой mail", record.Host)
			continue
		}

	}

	return conn, nil
}
