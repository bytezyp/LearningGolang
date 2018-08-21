package main

import "fmt"

func demo(number []int)  {
	a := []int{2,4}
	number = a
	//number = []int{2,4}
	fmt.Printf("%p, %v \n", number, number)
}
func main()  {
	var number []int
	demo(number)
	fmt.Printf("%p, %v", number, number)
}













