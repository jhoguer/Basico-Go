package main

import (
	"fmt"
	"net/http"
)

func HandlerRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World desde el handler")
}

func HandlerHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this is the API Endpoint")
}
