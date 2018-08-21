package main

import "fmt"

func main() {
	demo := make([]string, 0)
	a := []string{"1", "2"}
	b := []string{"1"}
	demo = append(demo,a...)
	fmt.Println(demo, append(demo, b...))
}




