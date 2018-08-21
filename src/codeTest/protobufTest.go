package main

import (
	"bytes"
	"fmt"
	"adinall"
	json2 "encoding/json"
)
func main()  {
	b := bytes.NewBufferString("zhang玉")
	fmt.Printf("%T", b, b.Bytes())
	adinall := proto.Unmarshal(body, req)
}
