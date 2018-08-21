package main

import (
	"fmt"
	"unsafe"
	"sync/atomic"
)

type aa struct {
	cc cc
}

type cc uintptr
func (c *cc) Check() {
	if uintptr(*c) != uintptr(unsafe.Pointer(c)) &&
		!atomic.CompareAndSwapUintptr((*uintptr)(c), 0, uintptr(unsafe.Pointer(c))) &&
		uintptr(*c) != uintptr(unsafe.Pointer(c)) {
		panic("sync.Cond is copied")
	}
}
func main()  {
	a := &aa{}
	b := a
	fmt.Println(a.c)
	fmt.Println(a.cc, b.cc)

}




