package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name     string
	LastName string
	Age      int
}

func mapToStruct(mp map[string]interface{}, iv interface{}) error {
	v := reflect.ValueOf(iv)
	if v.Kind() != reflect.Ptr {
		return fmt.Errorf("%T is not a pointer", iv)
	}

	v = v.Elem()
	if v.Kind() != reflect.Struct {
		return fmt.Errorf("%T is not a pointer to struct", iv)
	}

	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i) // reflect.StructField
		fv := v.Field(i)    // reflect.Value
		if val, ok := mp[field.Name]; ok {
			mfv := reflect.ValueOf(val)
			if mfv.Kind() != fv.Kind() {
				return fmt.Errorf("incompatible type %T for %q (%s)", val, field.Name, fv.Type())
			}
			if fv.CanSet() {
				fv.Set(mfv)
			}
		}
	}
	return nil
}

func main() {
	var st Student
	mp := map[string]interface{}{
		"Name":           "Mary",
		"Age":            42,
		"SomeMoreFields": true,
	}
	err := mapToStruct(mp, &st)
	fmt.Printf("STRUCT: %#v\nERR: %s\n", st, err)
}
