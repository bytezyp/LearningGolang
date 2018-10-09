package main

import (
	"strconv"
	"fmt"
	"os"
	"bufio"
	"log"
)

func main()  {

	dictFile, err := os.Open("delfile.txt")
	defer dictFile.Close()
	if err != nil {
		log.Fatalf("无法载入字典文件 \"%s\" \n")
	}
	reader := bufio.NewReader(dictFile)
	var text string
	var freqText string
	var frequency int
	var pos string
	fmt.Println(frequency, freqText)
	size, _ := fmt.Fscanln(reader, &text, &freqText, &pos)
	fmt.Println(size, text, freqText)
	frequency, err = strconv.Atoi(freqText)
	fmt.Println(frequency, err,size)

	dd := int(^uint(0) >> 1)
	fmt.Println(dd, uint(1))
}
