package main

import (
	"net/http"
	"fmt"
)


type Fu struct {
	Age int
	Name string
}

type Sun struct {
	Fu
	Age int
	Name string
	//FuInter
}

type FuInter interface {
	Say() int
}

type SunInter interface {
	http.ResponseWriter
	FuInter
	SaySun() int
}

func (s *Sun)SaySun() int  {
	return s.Age
}

func (f *Fu)Say() int  {
	return f.Age
}
func (s *Sun)Say() int  {
	return s.Age
}
func main() {
	testStruct := Sun{Fu{ 11,"Fu"},  22,"sun"}
	// 调用子类属性
	fmt.Println("Sun struct age :",testStruct.Age)
	// 调用父类属性
	fmt.Println("fu struct age :", testStruct.Fu.Age)
	// 调用子类方法
	fmt.Println("Sun say func :", testStruct.Say())
	// 调用父类方法
	fmt.Println("fu say func :", testStruct.Fu.Say())

	// 测试赋值变量
	testStruct2 := Sun{
		Age:22,
	}
	testStruct3 := Sun{Fu{ 11,"Fu"},  22,"sun"}
	fmt.Println(testStruct3, testStruct2)
	/*
	*	在给结构体赋值的时候，如果没有指定结构体的key，就必须全部给赋值。
	*	如果当时由于需求，只是想赋值指定字段，就要指定key，其他字段会golang会指定默认值。
	*
	*/

}