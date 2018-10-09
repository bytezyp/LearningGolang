package main

import (
	"fmt"
)

type Fu1 interface {
	Say()
}

type Persion struct {
	name string
	Fu1
}

func (p *Persion) Say() {
	fmt.Println("say...")
}
type Student struct {
	name string
}

func main()  {
	per := &Persion{}
	per.Say()
	//per.Fu1.Say()

}




