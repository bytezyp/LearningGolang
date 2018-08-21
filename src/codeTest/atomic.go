package main

import (
	"fmt"
	"sync/atomic"
)

func main()  {

	/*
	*	atomic包提供了底层原子级内存操作，对于同步算法的实现很有用。
	*	这些函数必须谨慎的保证正确的使用。除了某些特殊的底层应用，使用通道或者sync包的函数/类型实现同步更好。
	*
	*/

	var a int32=12
	fmt.Println(a)
	b := &a //在初始化变量b时， 类型被推断为指针类型。

	// atomic.Load* 系列函数 根据地址获取相应的值
	fmt.Println("atomic.LoadInt32(b)", atomic.LoadInt32(b))

	// atomic.Store* 系列函数，根据地址存储值，但没有返回值
	atomic.StoreInt32(b, 22)
	fmt.Println("atomic.StoreInt32(b, 22)", *b) // 获取其值

	// atomic.Add* 系列函数，根据地址的值，在原值的基础上，添加值，但有返回值。
	fmt.Println("atomic.AddInt32(b, 33)",atomic.AddInt32(b, 33))

	// atomic.Swap* 系列函数，在地址上，赋值上新值，返回旧的值，用于交换两个值
	var bb int32
	bb = 99
	fmt.Println("atomic.SwapInt32(bb, 33)", atomic.SwapInt32(&bb, 33))

	// atomic.CompareAndSwap* 系列函数，某个地址，是否是某个变量的地址，如果是，返回true，否则，返回false。
	fmt.Println("atomic.CompareAndSwapInt32(addr, old, new)", atomic.CompareAndSwapInt32(b,bb,66))
}











