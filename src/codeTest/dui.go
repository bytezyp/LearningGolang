package main

import (
	"fmt"
	"math"
)

func main()  {
	arr := [5]int{3,8,5,9,2}
	fmt.Println(len(arr))
	for i := 4; i >=0 ; i-- {
		num := i/2 - 1
		gen := int(math.Floor(float64(num)))
		l := 2*gen +1
		r := 2*gen + 2
		fmt.Println(l,r)
		//if arr[gen] < l && arr[gen] > arr[r] {
		//	arr[l] = arr[gen]
		//}
		//if arr[gen] < r && arr[gen] > arr[l] {
		//	arr[r] = arr[gen]
		//}else {
		//	arr[l], arr[r] = arr[r], arr[l]
		//}
	}

	//fmt.Println(arr)
}









