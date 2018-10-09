package packTest

import "fmt"

type P2 struct {
	P
	Name string
}

func (p *P2)Post()  {
	fmt.Println("p2 Testing", p.Name)
}

func init()  {
	fmt.Println("p2 init ing")
}
