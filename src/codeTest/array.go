package main

import (
	"fmt"
	"context"
	"time"
)

func main() {
	 array :=[...]int{1,3,4,5,5}
	var array1 [5]int
	var array2 =[5]int{1,2,3,4,5}
	var array3  = [...]int{1,2,3}
	array3[2] = 33
	fmt.Println(array, array1, array2, array3, len(array3), cap(array3))


	var a int
	a = 123
	b := float32(a)
	fmt.Println(a, b)
	context.WithTimeout(context.Background(), 5*time.Second)
	nowHour := time.Now().Unix() / 3600
	fmt.Println(nowHour, time.Now().Unix())
}



