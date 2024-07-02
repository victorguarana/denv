package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/path/{param}", func(w http.ResponseWriter, r *http.Request) {
		param := r.PathValue("param")
		response := fmt.Sprintf("Golang Response with param: %s\n", param)
		_, err := w.Write([]byte(response))
		if err != nil {
			http.Error(w, "Failed to write response", http.StatusInternalServerError)
		}
	})

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		panic(err)
	}
}
