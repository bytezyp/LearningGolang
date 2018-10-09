package main

import (
	"fmt"
	"reflect"
)

// 形状接口
type Shaper interface {
	Area() float32
}

//  正方形 类型
type Square struct {
	side float32
}

// 计算面积
func (sq *Square)Area() float32  {
	return sq.side * sq.side
}
// 长方形 类型
type Rectangle struct {
	length, width float32
}
// 计算面积
func(r Rectangle)Area() float32 {
	return r.length * r.width
}
func main()  {

	/*
	*
	*	接口基础用法
	*
	*/
	// 定义类型变量
	sql := new(Square)	// 引用类型变量
	//sql1 := Square{6} // 测试不是引用类型变量
	sql.side = 5
	// 定义接口变量
	var areaIntf  Shaper
	fmt.Println("初始化接口变量 reflect.type：", reflect.TypeOf(areaIntf))
	fmt.Println("初始化接口变量：reflect.value :", reflect.ValueOf(areaIntf))
	fmt.Println("初始化接口变量：value :", areaIntf)
	//areaIntf = sql1 测试非引用类型的变量，是否合法

	// 类型变量赋值给指定接口（相当于验证了类型变量是否实现接口）
	// gin 框架中，经常用 var _ Interface = &Type 来验证是否实现指定接口
	fmt.Printf("the Squarea has area: %f \n", areaIntf.Area())
	// 另一种方式
	areaIntf1 := Shaper(sql)
	fmt.Printf("the Squarea has area: %f \n", areaIntf1.Area())
	fmt.Println("接口变量赋值之后的状态 reflect.type：", reflect.TypeOf(areaIntf))
	fmt.Println("接口变量赋值之后的状态：reflect.value :", reflect.ValueOf(areaIntf))
	fmt.Println("接口变量赋值之后的状态：value :", areaIntf)

	/*
	*
	*	接口是方法的集合，但是在go中，它也可以是一个变量，由两部分组成类型和值，初始化的值是nil
	*	当接口变量，被赋值之后，接口变量就有了变量的引用，可以通过引用，调用类型变量的方法。
	*	注意：赋值给接口变量的值，必须引用类型的值，如果编译不过，不能进行转化。
	*/

	/*
	*
	*	接口用法1、 多态的变现形式
	*	在调用 shapes[n].Area() 这个时，只知道 shapes[n] 是一个 Shaper 对象，
	*	最后它摇身一变成为了一个 Square 或 Rectangle 对象，并且表现出了相对应的行为
	*/
	r := Rectangle{5,3}
	q := &Square{5}
	shapers := []Shaper{r, q}
	for n, shaper :=range shapers {
		fmt.Println("shape details:", shapers[n])
		fmt.Println("Area of this shape is:", shapers[n].Area())
		fmt.Println("anther Area of this shap is",shaper.Area())
	}
	/*
	*	这样就通过接口，这样可以通过接口，产生好多更干净，更简单的以及更有扩展性的代码。
	*
	*/
}









