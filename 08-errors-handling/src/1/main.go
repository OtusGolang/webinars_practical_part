package main

import (
	"errors"
	"fmt"
)

func foo() (bool, error) {
	return false, errors.New("foo failed")
}

func bar() (bool, error) {
	return false, fmt.Errorf("user %q not found", "test")
}

func main() {
	_, err := foo()
	if err != nil {
		fmt.Println(err)
	}

	if _, err := bar(); err != nil {
		fmt.Println(err)
	}
}
