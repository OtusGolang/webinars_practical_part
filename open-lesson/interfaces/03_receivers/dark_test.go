package receivers

import (
	"fmt"
	"unsafe"
)

func dark_magic(i any) {
	p := (*struct { // iface mock
		tab  unsafe.Pointer // itab mock
		data unsafe.Pointer
	})(unsafe.Pointer(&i))

	fmt.Printf("iface: {itab: %x, data: %x}\n", p.tab, p.data)
}
