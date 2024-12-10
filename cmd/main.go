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

	fmt.Println(*res)
}
