package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Path[1:]
	for i := 0; i < 10; i++ {
    	fmt.Fprintf(w, "He lives in a pineapple under the sea ... %s!\n", r.URL.Path[1:])
    	if name == "Spongebob Squarepants" {
    		fmt.Fprintf(w, "Absorbant and yellow and porous is he!\n")
    	}
    }
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}