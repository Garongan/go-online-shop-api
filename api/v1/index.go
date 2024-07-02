package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
	resp["message"] = "Hello World from Go! 👋"
	resp["language"] = "go"
	resp["cloud"] = "Hosted on Vercel! ▲"
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		fmt.Printf("Error happened in JSON marshal. Err: %s", err)
	} else {
		_, err := w.Write(jsonResp)
		if err != nil {
			return
		}
	}
}
