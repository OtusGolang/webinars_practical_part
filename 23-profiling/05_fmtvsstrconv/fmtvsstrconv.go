package fmtvsstrconv

import (
	"fmt"
	"strconv"
)

func Slow() string {
	return fmt.Sprintf("%d",42)
}

func Fast() string {
	return strconv.Itoa(42)
}
