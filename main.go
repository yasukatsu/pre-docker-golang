package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

type hoge struct {
	Status  int    `json:"status"`
	Results string `json:"results"`
}

func main() {
	fmt.Println("test")
	http.HandleFunc("/", rootHandler)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {

	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}
	fmt.Printf("%v\n", string(dump))

	w.Header().Set("Content-Type", "application/json")
}
