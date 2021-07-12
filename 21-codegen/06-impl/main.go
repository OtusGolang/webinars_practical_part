package main

import "io"

// r *myReader io.Reader
type myReader struct{}

var _ io.Reader = (*myReader)(nil)
