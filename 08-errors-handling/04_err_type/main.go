package main

import (
	"fmt"
	"os"
)

type PathError struct {
	Op   string // "open", "unlink", etc.
	Path string // The associated file.
	Err  error
}

func (e *PathError) Error() string {
	return e.Op + " " + e.Path + " " + e.Err.Error()
}

func checkConfig(path string) error {
	_, err := os.ReadFile(path)
	if err != nil {
		return &PathError{"open", path, err}
	}

	return nil
}

func main() {
	err := checkConfig("/etc/apt/sources.list")

	switch err := err.(type) { // prefer errors.As
	case *PathError:
		fmt.Println("path error: ", err)
	case nil:
		fmt.Println("success")
	default:
		fmt.Println("unknown error")
	}
}
