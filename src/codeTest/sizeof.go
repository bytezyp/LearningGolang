package main

import (
	"fmt"
	"unsafe"
)

func main()  {
	a := "90b452299f76883b0c0682db025e973f"

	fmt.Println(unsafe.Sizeof(a))
}






