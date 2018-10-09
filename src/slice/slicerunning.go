package main

import "fmt"

func Print1(p []int)  {
	fmt.Printf(" 切片-p: %p 切片-v: %v 切片长度：%v 切片容量：%v \n",p, p,len(p),cap(p))
}
/*
*
*	关于切片的一点总结
*/
func main()  {
	var ss []int
	ssV1 := make([]int, 0, 0)
	// []int 和make([]int, 0, 0) 效果一样
	Print1(ss)
	Print1(ssV1)
	// 切片len 一定要 < cap
	// 否则报错：len larger than cap in make([]int)
	ssv2 := make([]int, 4, 8)
	Print1(ssv2)
	fmt.Println("--------当切片容量和长度都是0的情况-------")
	sliceRunV1 := make([]int, 0)
	for i := 0; i < 10; i++{
		sliceRunV1 = append(sliceRunV1,i)
		Print1(sliceRunV1)
	}
	// 结论：切片长度和容量动态变化，不断重新分配内存，效率低下
	fmt.Println("--------当切片容量是10， 长度都是0的情况-------")
	sliceRunV2 := make([]int, 0, 10)
	for i := 0;i < 10; i++{
		sliceRunV2 = append(sliceRunV2,i)
		Print1(sliceRunV2)
	}
	// 结论：切片长度动态变化，容量不变，内存不从心分配，效率高

	fmt.Println("--------当切片容量是10， 长度都是10的情况-------")
	sliceRunV3 := make([]int, 10, 10)
	for i := 0;i < 10;i++{
		sliceRunV3 = append(sliceRunV3,i)
		Print1(sliceRunV3)
	}
	// 结论： 切片长度动态变量，容量是长度加起初容量，内存不从新分配，效率高，会出现默认值，不利于操作。
	fmt.Println("--------当切片容量是0， 长度都是0的情况, 指针用法-------")
	sliceRunV4 := make ([]int,0);
	slicepre:=&sliceRunV4;
	for i := 0;i < 10;i++{
		*slicepre=append(*slicepre, i)
		fmt.Printf("切片-p:%p \t指针-p:%p \t指针的值-v:%v\t 指针值的长度: %d \t指针值的容量：%d \n",sliceRunV4,slicepre, *slicepre, len(*slicepre), cap(*slicepre));
	}
	// 结论： 指针的效果和切片容量和长度为0时，差不多，指针的长度和容量是动态变化的。
}

