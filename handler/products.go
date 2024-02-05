package handler

import (
	"fmt"
	"net/http"
)

func Products(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from products")
}
