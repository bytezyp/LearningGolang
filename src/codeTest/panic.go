package main

import "fmt"

func demo()  {
	// 我的理解：panic 相当于 php中的 exit
	panic("panic ing exit")
	fmt.Println("this is a end ")
}


func main()  {
	demo()
}



