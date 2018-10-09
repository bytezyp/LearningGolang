package main

import (
	"fmt"
	"reflect"
)

type Inter interface {
	InterFunc1()
	InterFunc2()
}

type MyInt int
func main()  {

	var my MyInt = 78
	fmt.Println(reflect.TypeOf(my), reflect.ValueOf(my), reflect.ValueOf(my).Type(),
		reflect.ValueOf(my).Kind())
	v := reflect.ValueOf(my)
	y := v.Interface().(MyInt)
	fmt.Println(y, v.Interface(), v)
	//var rr io.Reader
	//tty, _ := os.Open("./fileTest.txt")
	//rr = tty

	//var x float64 = 3.4
	////var y int = 44
	//var d float64
	//d = 123
	//fmt.Println("type", reflect.TypeOf(x).Kind(), " value",reflect.ValueOf(x))
	//reflect.ValueOf(x).SetFloat(d)
	//fmt.Println(x)
}
