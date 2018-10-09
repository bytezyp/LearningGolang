package main

import (
	"fmt"
	"reflect"
	"unsafe"
	"time"
)

type data struct {}

func (this *data)Error() string {
	return "123"
}

func test() error {
	var p *data = nil
	return p
}

type T struct {

}

func main()  {
	p := new(interface{})

	print(p == nil, "---")
	print(*p == nil, "---")

	q := new(chan int)
	print(q == nil, "---")
	print(*q==nil)
	return
	var e error  = test()
	if e == nil {
		fmt.Println("e is nil")
	} else {
		fmt.Println("e is not nil")
	}

	fmt.Println(reflect.TypeOf(e))
	fmt.Println(reflect.TypeOf(e).Kind())
	fmt.Println(reflect.ValueOf(e))
	fmt.Println(reflect.ValueOf(e).Type())

	d := (*struct {
		itab uintptr
		data uintptr
	})(unsafe.Pointer(&e))
	fmt.Println(d)
	 err := test()
	 a := err.(*data)
	 
		fmt.Println(a, "----")


	if err := test(); err != nil {
		panic("333333")
	}
}







