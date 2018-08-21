package main

import "packTest"

func main()  {
	p := packTest.P{}
	p0 := packTest.P{}
	p1 := packTest.P1{p, "p1"}
	p1.Get()
	p2 := packTest.P2{p0,"p2"}
	p2.Post()
}


