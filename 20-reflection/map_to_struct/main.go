package main

import (
	"errors"
	"fmt"
	"reflect"
)

type St struct {
	Name     string
	LastName string
	Age      int
}

func mapToStruct(mp map[string]interface{}, iv interface{}) error {
	v := reflect.ValueOf(iv)
	if v.Kind() != reflect.Ptr {
		return errors.New("not a pointer to struct")
	}
	v = v.Elem()
	if v.Kind() != reflect.Struct {
		return errors.New("not a pointer to struct")
	}
	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i) // reflect.StructField
		fv := v.Field(i)    // reflect.Value
		if val, ok := mp[field.Name]; ok {
			mfv := reflect.ValueOf(val)
			if mfv.Kind() != fv.Kind() {
				return errors.New("incomatible type for " + field.Name)
			}
			fv.Set(mfv)
		}
	}
	return nil
}

func main() {
	var st St
	mp := map[string]interface{}{
		"Name":           "Mary",
		"Age":            42,
		"SomeMoreFields": true,
	}
	err := mapToStruct(mp, &st)
	fmt.Printf("STRUCT: %#v\nERR: %s\n", st, err)
}
