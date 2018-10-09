package main

import (
	"fmt"
	"reflect"
)

func main() {
	var interf interface{}
	fmt.Println(reflect.TypeOf(interf), reflect.ValueOf(interf))
	if interf == nil {
		fmt.Println(" interface is equal as nil")
	} else {
		fmt.Println("interface is not equal as nil")
	}
	interf = 56
	fmt.Println(reflect.TypeOf(interf), reflect.ValueOf(interf))
	if interf == nil {
		fmt.Println(" interface is equal as nil")
	} else {
		fmt.Println("interface is not equal as nil")
	}

	var aa int  = 123
	//fmt.Println("aaa",reflect.ValueOf(aa).Interface().(int), reflect.ValueOf(aa).Interface())
	v := reflect.ValueOf(&aa)
	//v.SetInt(33)
	fmt.Println(v.CanSet())

	var f float64 = 3.4
	p := reflect.ValueOf(&f)
	t := reflect.TypeOf(f)
	fmt.Println(p, t.Method)
	//pv := p.Elem()
	//fmt.Println(p.CanSet(), p, pv)
}
