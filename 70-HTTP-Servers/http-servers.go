package main

import (
	"fmt"
	"net/http"
)

type handler struct {
}

func (h handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "welcom\n")
}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v :%v\n", name, h)
		}
	}
}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	//默认请求
	http.ListenAndServe(":8090", handler{})

}
