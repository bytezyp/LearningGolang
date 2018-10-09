package main

import (
	"unsafe"
	"fmt"
	"hash/crc32"
)

type data struct {}

func (this *data) Error() string {
	return ""
}

func test() error {
	var p *data = nil
	return p
}
type Hash func(data []byte) uint32

type Map struct {
	hash     Hash
	replicas int
	keys     []int // Sorted
	hashMap  map[int]string
}
func New(replicas int, fn Hash) *Map {
	m := &Map{
		replicas: replicas,
		hash:     fn,
		hashMap:  make(map[int]string),
	}
	fmt.Println(fn)
	if m.hash == nil {
		m.hash = crc32.ChecksumIEEE
	}
	return m
}
func main()  {
	a := New(2, nil)
	fmt.Println(a.hash)
	var e error = test()
	d := (*struct {
		itab uintptr
		data uintptr
	})(unsafe.Pointer(&e))
	fmt.Println(d)
	var array = [...]int{1,2,3}
	fmt.Println(len(array),cap(array))


}


