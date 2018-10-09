package main

import "net/http"

type res struct {

}

type Rest interface {
	http.ResponseWriter
}

var _ Rest = &res{}

func main()  {

}












