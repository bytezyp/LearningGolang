package main

import (
	"fmt"
	"reflect"
)

// 定义两个类型stockPostion 和car，一个接口valuable,包含一个方法 getValue
// 停车位
type StockPostion struct {
	ticker string
	sharePrice float32
	count float32
}

func (s *StockPostion)getValue() float32 {
	return s.count * s.sharePrice
}

type car struct {
	make string
	model string
	price float32
}

func (c *car)getValue() float32 {
	return c.price
}

type valuable interface {
	getValue() float32
}

func showValue(value valuable)  {
	fmt.Println("value is : ", value.getValue(), " ，变量类型：", reflect.TypeOf(value), "，可以看出类型变成引用类型")
}

func main()  {
	// 刚才犯了一个错误，忘了加 & 取地址符，编译出现问题，
	ss := &StockPostion{"test", 2, 2}
	cc := &car{"china", "smart", 33}
	showValue(ss)
	fmt.Println("--------")
	showValue(cc)

	var num int = 33
	fmt.Println(reflect.TypeOf(num), "标准类型的type")
}






