package main

import (
	"math"
	"fmt"
	"github.com/willf/bloom"
	"github.com/willf/bitset"
)

func main()  {
	log := math.Log(2)
	pow := math.Pow(log,2)
	fmt.Println(log, pow)
	a := math.Pow(math.Log(2), 2)
	b := -1 * float64(10) * math.Log(0.12) / math.Pow(math.Log(2), 2)
	fmt.Println(b, math.Log(0.0012))
	bb,tt := bloom.EstimateParameters(32,0.0001)
	fmt.Println(bb, tt, a) // 4294967296
	fmt.Println(bitset.Cap())
}
