package main

import (
	"fmt"
	"unsafe"
)

func main() {
	userSize()
	padding()
	pointers()
	slicePtr()
	dereference()
}

func userSize() {
	var s struct{}
	fmt.Printf("%+v %d\n", s, unsafe.Sizeof(s))

	var u0 User // Zero Value для типа User
	fmt.Printf("%+v %d\n", u0, unsafe.Sizeof(u0))

	u1 := User{} // Zero Value для типа User
	fmt.Printf("%+v %d\n", u1, unsafe.Sizeof(u1))

	u2 := &User{} // То же, но указатель
	fmt.Printf("%+v %d\n", u2, unsafe.Sizeof(u2))

	u3 := User{1, "Vasya", 23, nil} // По номерам полей
	fmt.Printf("%+v %d\n", u3, unsafe.Sizeof(u3))

	u4 := User{ // По именам полей
		Id:      1,
		Name:    "Vasya",
		friends: []int64{1, 2, 3},
	}
	fmt.Printf("%+v %d\n", u4, unsafe.Sizeof(u4))
}

func padding() {
	fmt.Println(unsafe.Sizeof(1))   // 4 на моей машине
	fmt.Println(unsafe.Sizeof("A")) // 8 (длина + указатель)

	var x struct {
		a bool   // 1
		b bool   // 1
		c string // 8
	}
	// 0   1       8
	// a   b       c
	// x

	fmt.Println(unsafe.Sizeof(x)) // 12!
	fmt.Println(
		unsafe.Offsetof(x.a), // 0
		unsafe.Offsetof(x.b), // 1
		unsafe.Offsetof(x.c)) // 4
}

func pointers() {
	x := 1
	fmt.Printf("%v %T\n", x, x)

	xPtr := &x
	fmt.Printf("%v %T\n", xPtr, xPtr)

	var p *int
	fmt.Printf("%v %T\n", p, p)
	fmt.Printf("%v %T\n", *p, p)
}

func slicePtr() {
	s := []int{1, 2, 3}
	p := &s[1]
	fmt.Println(*p, s) // 2 [1, 2, 3]

	*p = 10
	fmt.Println(*p, s) // ?

	s = append(s, 4)
	*p = 20
	fmt.Println(*p, s) // ?
}

func dereference() {
	type vertex struct {
		x, y int
	}

	p := vertex{1, 3}
	pPtr := &p

	fmt.Println(pPtr.x) // (*pPtr).x
	fmt.Println((*pPtr).y)
}
