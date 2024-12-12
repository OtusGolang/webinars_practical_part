package main

import (
	"errors"
	"fmt"
)

var Err42 = errors.New("oh no, it is the 42 error")

func main() {
	err := doManyRollbacks()
	if err != nil {
		fmt.Println("Error: ", err)
	}
	if errors.Is(err, Err42) {
		fmt.Println("42 is in there! 'Is' works for multierrors!")
	}

}

func doManyRollbacks() error {
	var mErr []error
	for _, rollbackReq := range prepareRollbackRequests() {
		mErr = append(mErr, rollback(rollbackReq))
	}
	return fmt.Errorf("mass rollback: %w", errors.Join(mErr...))
}

func prepareRollbackRequests() []interface{} {
	return []interface{}{42, 43, 44}
}

func rollback(v interface{}) error {
	// do something...

	if v == 42 {
		return Err42
	}

	return fmt.Errorf("uh oh %v", v) // try this
}
