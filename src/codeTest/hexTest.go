package main

import (
	"encoding/hex"
	"fmt"
	"bytes"
)

func main()  {
	// 将[]byte 转化成16进制的字符串
	src := []byte("hello")
	encodeStr := hex.EncodeToString(src)
	fmt.Println(encodeStr)

	// 16进制的字符串转化为[]byte
	test,_ := hex.DecodeString(encodeStr)
	fmt.Println(bytes.Compare(src,test))

	b := make([]byte, 1<< 16)
	fmt.Println( len(b))

}


