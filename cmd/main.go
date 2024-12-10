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

	syntax := verifier.ParseAddress(email)
	if !syntax.Valid {
		fmt.Println("Некорректный email")
	}

	fmt.Println("username:", syntax.Username)
	fmt.Println("domain:", syntax.Domain)
}
