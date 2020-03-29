package stringbuilder

import "strings"

func Slow() string {
	a := ""
	a += "1"
	a += "2"
	a += "3"
	a += "4"
	a += "5"
	return a
}

func Fast() string {
	builder := strings.Builder{}
	builder.WriteString("1")
	builder.WriteString("2")
	builder.WriteString("3")
	builder.WriteString("4")
	builder.WriteString("5")
	return builder.String()
}

func VeryFast() string {
	return "1" + "2" + "3" + "4" + "5"
}