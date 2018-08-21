package main

import (
	"math"
	"fmt"
)

const abortIndex int8 = math.MaxInt8 / 2 // 防止++时反转，取最大值的一半
func main() {
	fmt.Println(abortIndex)
}
