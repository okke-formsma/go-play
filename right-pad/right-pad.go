package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func handler(w http.ResponseWriter, r *http.Request) {
	str := r.URL.Query().Get("str")
	ch := r.URL.Query().Get("ch")
	if ch == "" {
		ch = " "
	}
	length, _ := strconv.Atoi(r.URL.Query().Get("len")) //length zero in case of errors
	if length > 1024 {
		fmt.Fprintf(w, "Len exceeds 1024 characters. Please contact sales for an enterprise license.")
		return
	}

	// This is where the magic happens. The inner sanctum. The inner loop.
	for len(str) < length {
		str = str + ch
	}

	fmt.Print(str)
	res, err := json.Marshal(struct {
		string `json:"str"`
	}{str})
	if err != nil {
		fmt.Fprintf(w, "unable to encode '%s'", str)
	}

	fmt.Fprintf(w, "%s", res)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
