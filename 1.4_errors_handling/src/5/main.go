package main

import (
	"fmt"

	"github.com/pkg/errors"
)

const testEmail = "test@mail.ru"

func getUserEmail(id int) (string, error) {
	if id == 1 {
		return testEmail, nil
	}

	return "", errors.New("email not found")
}

func loginUserById(id int) (bool, error) {
	email, err := getUserEmail(id)
	if err != nil {
		return false, errors.Wrap(err, "unable to get email")
	}

	if email != testEmail {
		return false, errors.New("unable to login user")
	}

	return true, nil
}

func main() {
	if _, err := loginUserById(0); err != nil {
		fmt.Printf("%+v\n", err)
	} else {
		fmt.Println("success")
	}
}
