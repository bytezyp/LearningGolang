package main

import "fmt"

func ret() (i int) {
	defer func() {i++}()
	return 1
}
func main()  {
	fmt.Println(ret())
}
