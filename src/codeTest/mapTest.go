package main

import (
	"os"
	"bufio"
	"io"
	"fmt"
)
type StrMap map[string]bool
func main()  {
	aa := make(StrMap)
	file,_ := os.Open("fileTest.txt")
	inputReader := bufio.NewReader(file)
	//delFile,_ := os.Create("delfile.txt")
	for i := 1 ; i < 100000; i++ {
		inputString, readerError := inputReader.ReadString('\n')
		if readerError == io.EOF {
			return
		}
		if aa[inputString]  {
			fmt.Println(aa[inputString])
		} else {
			//fmt.Println(aa[inputString])
			aa[inputString] = true
		}
		fmt.Println(aa[inputString])
	}
}
