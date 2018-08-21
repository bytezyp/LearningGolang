package main

import (
	"runtime"
	"fmt"
)

func main()  {
	b := make([]byte, 1<<16)
	b = b[:runtime.Stack(b, true)]
	fmt.Println(b)
	fmt.Printf("%s", b)
}
