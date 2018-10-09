package main

import "fmt"

func main()  {
	var ss  = make([]int, 0)
	for i := 1; i< 10 ; i++ {
		ss = append(ss, i)
	}
	fmt.Println(ss)
	// 目标长度 >=  源长度
	// 目标切边len:10, 源数据len：10
	// 结论： 会被全部复制
	//ssCopyV := make([]int, 10)
	ssCopyV := make([]int, 13)
	num := copy(ssCopyV, ss)
	fmt.Println(num, ssCopyV)
	// 目标长度 < 源长度
	// 目标切边len:10, 源数据len：10
	// 结论： 只复制目标长度的值
	ssCopyV2 := make([]int, 5)
	num = copy(ssCopyV2, ss)
	fmt.Println(num, ssCopyV2)
	// 目标长度 == 0
	// 目标切边len:10, 源数据len：10
	// 结论： 不能复制，但是不报错
	ssCopyV3 := make([]int, 0)
	num = copy(ssCopyV3, ss)
	fmt.Println(num, ssCopyV3)

	/*
	*	总结： 切片复制，目标切片长度一定要 >= 源切片的长度。
	*	初始化切片时，切片的长度，必须大于0
	*/
}



