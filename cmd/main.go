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

	verifier := verifier.NewVierifier()

	client, err := verifier.Verify(email)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Email:", email)
	fmt.Println("Client:", client)

}
