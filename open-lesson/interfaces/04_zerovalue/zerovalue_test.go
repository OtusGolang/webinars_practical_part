package zerovalue

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func ReadFile(fname string) error {
	var err *os.PathError // nil

	if err == nil {
		log.Println("err is nil")
	}

	// Do some work without err...
	return nil
}

func TestReadFile(t *testing.T) {
	if err := ReadFile(""); err != nil {
		log.Printf("ERR: (%T, %v)", err, err)
	} else {
		log.Println("OK")
	}
}

func TestZeroValues(t *testing.T) {

	var data *byte
	var in interface{}

	fmt.Println("1: ", data, data == nil) //prints: <nil> true
	fmt.Println("2: ", in, in == nil)     //prints: <nil> true

	in = data
	fmt.Println("3: ", in, in == nil) //prints: <nil> false
	//'data' is 'nil', but 'in' is not 'nil'

}
