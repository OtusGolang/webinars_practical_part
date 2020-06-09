package main

import (
	"fmt"
)

type countriesCodes map[string]int

type world struct {
	Name      string
	countries countriesCodes
}

func NewWorld(name string) *world {
	w := world{
		Name:      name,
		countries: make(countriesCodes),
	}
	return &w
}

func (w *world) AddCountry(name string, code int) {
	w.countries[name] = code
}

func (w world) GetCountryCode(name string) (code int, ok bool) {
	code, ok = w.countries[name]
	return
}

func main() {
	w := NewWorld("Earth")
	w.AddCountry("Russia", 666)
	c, _ := w.GetCountryCode("Russia")
	fmt.Println(c) // 666
}
