package handler

import (
	"fmt"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Hello From Go")
	if err != nil {
		return
	}
}
