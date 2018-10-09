package packTest

import "fmt"

type P1 struct {
	P
	Name string
}


func (p *P1)Get()  {
	fmt.Println("p1 Testing", p.Name)
}

func init()  {
	fmt.Println("p1 init ing")
}