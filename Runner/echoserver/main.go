package main

import (
	"fmt"
	"net/http"
)

func handleRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Fprintf(w, "%+v\n", r)
	default:
		fmt.Fprintf(w, "Sorry, only GET methods are supported.")
	}
}

func main() {
	http.HandleFunc("/", handleRequest)
	if err := http.ListenAndServe(":58080", nil); err != nil {
		panic(err)
	}
}
