package main

import (
	"os"
	"os/signal"
	"syscall"
	"runtime"
	"fmt"
	"time"
)

func main()  {
	setupSigusr1Trap()

	go a()
	m1()
}

func m1()  {
	m2()
}

func m2()  {
	m3()
}

func m3()  {
	time.Sleep(time.Hour)
}

func a()  {
	time.Sleep(time.Hour)
}

func setupSigusr1Trap()  {
	c := make(chan os.Signal, 1)
	fmt.Println(len(c))
	go func() {
		//for  range c {
			DumpStacks()
		//}
	}()
	signal.Notify(c, syscall.SIGUSR1)
	fmt.Println(len(c))

}

func DumpStacks()  {
	buf := make([]byte, 1<<16)
	buf = buf[:runtime.Stack(buf, true)]
	fmt.Println(buf)
	fmt.Printf("=== BEGIN goroutine stack dump ===\n%s\n=== END goroutine stack dump ===", buf)
}













