package stringsvsre

import (
	"regexp"
	"strings"
)

var (
	re *regexp.Regexp
)

func Fast() bool {
	return strings.Contains("Hello world of golang", "world")
}

func VerySlow() bool {
	 res, _ := regexp.MatchString("world", "Hello world of golang")
	 return res
}

func Slow() bool {
	return re.MatchString("Hello world of golang")
}

func init() {
	re, _ = regexp.Compile("world")
}