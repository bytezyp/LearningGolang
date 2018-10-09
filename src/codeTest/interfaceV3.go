package main

import "fmt"

type OneInter interface {
	Onesay()
}

type TwoInter interface {
	Twosay()
}

type ThreeInter interface {
	OneInter
	TwoInter
	eat()
}
type test struct {
	name string
	ThreeInter
}

func (t *test) Onesay() {
	fmt.Println("i am onesay")
}

func (t *test)Twosay()  {
	fmt.Println("i am twosay")
}
//func (t *test)eat()  {
//	fmt.Println("i am eat")
//}
var _ ThreeInter = &test{}
func main()  {

}








