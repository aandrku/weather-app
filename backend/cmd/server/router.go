package main

import "net/http"

func router() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /weather", nil)

	return mux
}
