package main

import (
	"fmt"

	"github.com/pkg/errors"
)

const testEmail = "test@mail.ru"

type UserNotFound struct{}

func (e UserNotFound) Error() string {
	return "user not found"
}

func getUserEmail(id int) (string, error) {
	if id == 1 {
		return testEmail, nil
	}

	return "", UserNotFound{}
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
	_, err := loginUserById(0)

	switch err := errors.Cause(err).(type) {
	case UserNotFound:
		fmt.Println("unable to find user: ", err)
	case nil:
		fmt.Println("success")
	default:
		fmt.Println("unkown error")
	}
}
