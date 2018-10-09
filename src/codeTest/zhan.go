package main

import (
	"fmt"
	"time"
)

func zhanFun1()  {
	time.Sleep(5 * time.Second)
	fmt.Println("zhanFun1.....")
}


func zhanFun() func() {
	return func() {
		fmt.Println("zhanFun start ing")
		zhanFun1()
		fmt.Println("zhanFun end ing")
	}
}

func main()  {
	zhanFun()()
}













