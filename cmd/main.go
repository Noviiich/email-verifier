package main

import (
	"fmt"
	"os"

	"github.com/Noviiich/email-verifier/pkg/verifier"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <email>")
		return
	}

	email := os.Args[1]

	res, err := verifier.VerifyEmail(email)
	if err != nil {
		fmt.Println(err)
		return
	}

	domain := *&res.Syntax.Domain

	mx, err := verifier.CheckMx(domain)
	if err != nil {
		fmt.Println("Нет такой mx-записи")
		return
	}

	for _, record := range mx.Records {
		fmt.Printf("Сервер: %s, Приоритет: %d\n", record.Host, record.Pref)
	}
}
