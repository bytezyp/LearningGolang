package main

import "fmt"

func Print(PP interface{})  {
	fmt.Printf("参数值：%v, 参数地址：%p \n", PP, PP)
}
func main()  {
	var ssV []string
	fmt.Printf("length:%v \naddr:%p \nisNil:%v\n", len(ssV), ssV, ssV == nil)
	fmt.Println("-----------------")
	// 切片追加(1):单个元素
	ssV3 := append(ssV, "22")
	ssV = append(ssV, "22")
	fmt.Printf("ssV: addr:%p , ssV3: addr:%p \n", ssV, ssV3)
	fmt.Println(ssV3)
	// 切片追加(2): 多个元素
	ssV = append(ssV, "33", "44")
	fmt.Printf("ssV-v:%v ssV-p:%p \n",ssV, ssV)
	// 切片追加(3): 单个切片元素, 或者说两个切片合并
	ssV = append(ssV, ssV3...)
	fmt.Printf("ssV-v:%v ssV-p:%p \n", ssV, ssV)
	// 切片删除(1): 删除指定元素，必须知道元素的下标, 例如：删除元素44 ，下表是 2
	index := 2
	delRes := append(ssV[:index],ssV[index+1:]...)
	ssV = append(ssV[:index],ssV[index+1:]...)
	Print(delRes)
	Print(ssV)
	// 切片中间插入元素：指定插入的位置, 55 例如：插入位置， index 1
	/*
	*	做法： 先生成一个临时的切片，然后把插入位置之前的元素放到临时切片中，
	*	然后在临时后，追加要插入的元素，最后把插入位置之后的元素追加的临时切片的后面。
	*/
	index = 1
	ssV1 := append([]string{}, ssV[:index]...)
	ssV1 = append(ssV1, "55")
	ssV = append(ssV1, ssV[index:]...)
	Print(ssV)

}






