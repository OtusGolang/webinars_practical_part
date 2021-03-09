package main

import (
	"fmt"
	"github.com/OtusGolang/webinars_practical_part/04-simple-data-types/example"
	"github.com/OtusGolang/webinars_practical_part/04-simple-data-types/runes"
)

var i float32 = 42

func notNakedReturn() (a int, b bool) {
	return
}

func main() {
	j := example.Answer()
	fmt.Println(j)
	fmt.Println(1 - 0.8 - 0.2)

	fmt.Println()
	runes.Example()
}
