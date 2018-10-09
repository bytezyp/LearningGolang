package main

import (
	"net/http"
	"time"
	"fmt"
)

func hello(wr http.ResponseWriter, r *http.Request)  {
	wr.Write([]byte("hello"))
}

func timeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(wr http.ResponseWriter, r *http.Request) {
		timeStart := time.Now()
		next.ServeHTTP(wr, r)
		timeElapsed := time.Since(timeStart)
		fmt.Println(timeElapsed)
	})
}


func main() {
	http.HandleFunc("/", timeMiddleware(hello))
}




