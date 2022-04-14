package main

import (
	"fmt"
	"net/http"
)

const MESSAGE = "hello world"
const ADDRESS = ":1024"

func main() {
	fmt.Println("Server listening...")
	http.HandleFunc("/hello", Hello)
	if e := http.ListenAndServe(ADDRESS, nil); e != nil {
		fmt.Println(e)
	}
}

func Hello(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content/Type", "text/plain")
	fmt.Fprintf(resp, MESSAGE)
}
