package main

import (
	"sync"
	"time"
)

var m *sync.RWMutex

func main() {
	m = new(sync.RWMutex)

	// 多个同时读
	go write(1)
	go write(2)

	// 多个同时读
	//go read(3)
	//go read(4)
	time.Sleep(6*time.Second)
}

func read(i int) {

	m.RLock()
	println(i,"read start")
	println(i,"reading")
	time.Sleep(1*time.Second)
	println(i,"read over")
	m.RUnlock()

}


func write(i int) {

	m.Lock()
	println(i,"write start")
	println(i,"writing")
	time.Sleep(1*time.Second)
	println(i,"write over")
	m.Unlock()

}