package main

import (
	"fmt"
	"reflect"
)

type RouterString func() string
type RouterBytes func() []byte
func AddRouter(router interface{}) {
	switch router := router.(type) {
	case RouterString:
		println("string",router)
	case RouterBytes:
		println("bytes", router)
	default:
		println("unknown types", router)
	}
}
func main() {
	var val interface{} = int64(58)
	fmt.Println(reflect.TypeOf(val))
	val = 50
	fmt.Println(reflect.TypeOf(val))
	var aa interface{} = 64
	if aa == nil {
		fmt.Println("aa = nil", reflect.TypeOf(aa), reflect.ValueOf(aa))
	} else {
		fmt.Println("aa != nil", reflect.TypeOf(aa), reflect.ValueOf(aa))
	}
	var cc interface{}

	var dd  = int(22)
	//var bb interface{} = &dd
	cc = &dd
	fmt.Println("cc", reflect.TypeOf(cc), reflect.ValueOf(cc))
	fmt.Println(reflect.TypeOf(aa), reflect.ValueOf(aa))
}
