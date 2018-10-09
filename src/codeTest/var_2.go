package main

import (

	"fmt"

)

var p *int

func foo() (*int, error) {

	var i int = 5

	return &i, nil

}

func bar() {

	//use p
	// 使用全局变量
	fmt.Println(*p)

}

func main() {

	fmt.Println(p)
	// 赋值给全局变量
	p, _ = foo()
	fmt.Println(p)


	bar()

	//fmt.Println(*p)

}
