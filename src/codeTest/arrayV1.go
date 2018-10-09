package main

import "fmt"

func main()  {
	var aa [5]int
	bb := make([]int, 5 ,5)
	bb[0] =1
	bb[1] =1
	bb[2] =1
	bb[3] =1
	bb[4] =1
	ee := make([]int, 0 ,11)
	ee = append(ee, bb...)
	ee =append(ee, 66)
	fmt.Println(len(aa), len(bb), aa, bb,bb[3:], bb[0:3], len(ee),cap(ee), ee)
}


