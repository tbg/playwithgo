package main

import (
	"fmt"
	"net/http"
)

type String string

func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, string(s))
}

func main() {
	// your http.Handle calls here
	http.ListenAndServe("localhost:4000", String("What up"))
}
