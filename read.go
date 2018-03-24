package main

import (
	"net/http"
)

func read(w http.ResponseWriter, r *http.Request) {
	b, ok := fsload(r.URL.Path)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	_, err := w.Write(b)
	if err != nil {
		lg.Printf("%s %v", r.URL.Path, err)
	}
}
