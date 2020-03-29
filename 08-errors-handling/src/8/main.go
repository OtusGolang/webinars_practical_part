package main

import (
	"errors"
	"fmt"
)

const testEmail = "test@mail.ru"

type UserNotFound struct {
}

func (e *UserNotFound) Error() string {
	return "user not found"
}

func getUserEmail(id int) (string, error) {
	if id == 1 {
		return testEmail, nil
	}

	return "", &UserNotFound{}
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
	_, err := loginUserById(0)

	switch err := errors.Unwrap(err).(type) {
	case *UserNotFound:
		fmt.Println("unable to find user: ", err)
	case nil:
		fmt.Println("success")
	default:
		fmt.Println("unknown error")
	}
}
