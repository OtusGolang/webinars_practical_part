package bench_test

import (
	"regexp"
	"testing"
)

var (
	re *regexp.Regexp
)

func match(str, match string) bool {
	res, _ := regexp.MatchString(match, str)
	return res

	// return re.MatchString(str)

	// return strings.Contains(str, match)
}

func BenchmarkMatch(b *testing.B) {
	re = regexp.MustCompile("world")

	for i := 0; i < b.N; i++ {
		match("Hello world of golang", "world")
	}
}
