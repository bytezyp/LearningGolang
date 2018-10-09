package main

type functest struct {
	name string
}

func main() {
	aa := functest{}
	aa.name = func() string {
		return "ddd"
	}()
}



