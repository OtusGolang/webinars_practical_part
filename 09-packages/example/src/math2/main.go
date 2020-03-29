package main

import (
	"fmt"

	c "calc"
	c2 "calc/internal"
)

func main() {
	fmt.Printf("Hello calc: %f\n", c.Antisin(3.1))
	fmt.Printf("Hello calc: %f\n", c2.Sin(3.1))
}
