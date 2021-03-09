package runes

import "fmt"

func Example() {
	//s := "Jack"
	//s := "Джек"
	s := "Дже♫к"
	for i, val := range s {
		fmt.Printf("%d %s %d\n", i, string(val), val)
	}
	fmt.Println("LEN: ", len(s))
}
