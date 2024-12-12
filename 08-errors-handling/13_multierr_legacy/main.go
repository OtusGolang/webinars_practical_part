package main

import (
	"fmt"
	"github.com/hashicorp/go-multierror"
)

func main() {
	var mErr error
	for _, rollbackReq := range prepareRollbackRequests() {
		mErr = multierror.Append(mErr, rollback(rollbackReq))
	}
	fmt.Println(mErr == nil)
}

func prepareRollbackRequests() []interface{} {
	return []interface{}{42}
}

func rollback(interface{}) error {
	return nil
}
