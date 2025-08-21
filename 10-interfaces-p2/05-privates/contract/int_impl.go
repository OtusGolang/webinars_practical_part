package contract

import "fmt"

type ContractInternalImpl struct {
}

func (c *ContractInternalImpl) Sign() error {
	return nil
}

func (c *ContractInternalImpl) cancel() error {
	fmt.Println("Nobody can implement cancel outside the package")
	return nil
}
