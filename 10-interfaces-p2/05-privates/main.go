package main

import (
	"fmt"

	"github.com/OtusGolang/webinars_practical_part/10-interfaces-p2/05-privates/contract"
	"github.com/OtusGolang/webinars_practical_part/10-interfaces-p2/05-privates/impl"
)

func main() {
	var c contract.Contract
	c.Sign()
	// c.cancel() // This will not compile because cancel is unexported

	// Can use internalImpl as contract.Contract
	internalImpl := &contract.ContractInternalImpl{}
	c = internalImpl
	// but internalImpl.cancel() or c.cancel() is not accessible here

	// Can not use externalImpl as contract.Contract
	externalImpl := &impl.ContractExternalImpl{}
	// c = impl  // Compilation error: cannot use impl (variable of type *impl.ContractExternalImpl) as contract.Contract value in assignment: *impl.ContractExternalImpl does not implement contract.Contract (unexported method cancel)

	externalImpl.Sign()
	// c.cancel() // Compilation error: c.cancel undefined (cannot refer to unexported method cancel)

	fmt.Println("Contract methods called successfully.")

}
