package main

import (
	"fmt"
	"sync"
)

type DD struct {
	Name string
	sync.Map
}

func main()  {
	d := &DD{Name:"zhang"}
	fmt.Println(d.count)
}






