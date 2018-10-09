package main

import (
	"fmt"
	"time"
)

func formatUnixTime()  {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
}

func main()  {
	formatUnixTime()
}

