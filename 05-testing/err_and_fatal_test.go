package main

import (
	"fmt"
	"os"
	"strconv"
	"testing"
)

func TestAtoi(t *testing.T) {
	os.Getwd()
	const str, want = "43a", 42
	got, err := strconv.Atoi(str)
	if err != nil {
		t.Fatalf("strconv.Atoi(%q) returns unexpeted error: %v", str, err)
	}
	if got != want {
		t.Errorf("strconv.Atoi(%q) = %v; want %v", str, got, want)
	}

	fmt.Printf("The end of the test")
}
