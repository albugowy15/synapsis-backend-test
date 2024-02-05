package handler

import (
	"fmt"
	"html"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Register %q", html.EscapeString(r.URL.Path))
}

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Login %q", html.EscapeString(r.URL.Path))
}
