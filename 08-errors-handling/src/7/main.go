package main

import (
	"fmt"
)

const testEmail = "test@mail.ru"

func getUserEmail(id int) (string, error) {
	if id == 1 {
		return testEmail, nil
	}

	return "", fmt.Errorf("email not found")
}

func loginUserById(id int) (bool, error) {
	email, err := getUserEmail(id)
	if err != nil {
		return false, fmt.Errorf("unable to get email: %w", err)
	}

	if email != testEmail {
		return false, fmt.Errorf("unable to login user")
	}

	return true, nil
}

func main() {
	if _, err := loginUserById(123); err != nil {
		fmt.Printf("%+v\n", err)
	} else {
		fmt.Println("success")
	}
}
