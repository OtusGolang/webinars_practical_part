package main

import (
	"reflect"
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestSetFail(t *testing.T) {
	v := 5.5
	rVal := reflect.ValueOf(v)
	assert.Panic(t, func() {
		rVal.SetFloat(1.1)
	}, "")

	t.Log(v)
}

func TestSetOk(t *testing.T) {
	v := 5.5
	rVal := reflect.ValueOf(&v)
	t.Log(rVal.String(), rVal)

	rVal.Elem().SetFloat(1.1)

	t.Log(v)
}
