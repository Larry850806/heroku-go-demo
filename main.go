package main

import (
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write([]byte(`<h1>Hello World. I'm Larry</h1>`))
	})

	http.ListenAndServe("0.0.0.0:"+os.Getenv("PORT"), nil)
}
