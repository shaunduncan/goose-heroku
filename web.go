package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", root)

	port := ":" + os.Getenv("PORT")
	if err := http.ListenAndServe(port, nil); err != nil {
		panic(err)
	}
}

func root(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprint(resp, "Hello, world")
}
