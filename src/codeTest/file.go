package main

import (
	"os"
	"fmt"
	"math/rand"
	"time"
)

func GetRandomString(leng int) string {
	allStr := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(allStr)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	//fmt.Println(r, "----", r.Intn(len(bytes)), "----",bytes)
	for i := 0; i < leng; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

type strMap map[string]int

func main()  {
	//fmt.Println(GetRandomString(32))
	file, err := os.Create("fileTest.txt")
	if err != nil {

	}
	for i := 1; i < 1000000; i++ {
		_,err := file.Write([]byte(GetRandomString(32)+"\n"))

		if err != nil {
			fmt.Println("add err")
		}
	}
	return
	//file,_ := os.Open("fileTest.txt")
	//inputReader := bufio.NewReader(file)
	//for i := 1 ; i < 10; i++ {
	//	inputString, readerError := inputReader.ReadString('\n')
	//	fmt.Printf("The input was: %s", inputString)
	//	if readerError == io.EOF {
	//		return
	//	}
	//}
}
