package main

import (
	"io"
	"net/http"
)

func write(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	p := make([]byte, *maxfs+1)
	n, err := r.Body.Read(p)
	if err != nil && err != io.EOF {
		lg.Printf("%s %v", r.URL.Path, err)
		http.Error(w, "Failed to read request body.", http.StatusInternalServerError)
		return
	}
	if n > *maxfs {
		http.Error(w, "Maximum file size exceeded.", http.StatusBadRequest)
		return
	}
	p = append([]byte(nil), p[:len(p)]...)
	fsstore(r.URL.Path, p)
}
