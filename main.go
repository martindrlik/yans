package main

import (
	"flag"
	"net/http"
)

var (
	addr  = flag.String("addr", ":8080", "listen on the TCP network address addr")
	maxfs = flag.Int("maxfs", 500, "maximum file size")
)

func main() {
	flag.Parse()
	http.HandleFunc("/write/", fixPath(len("/write"), write))
	http.HandleFunc("/read/", fixPath(len("/read"), read))
	http.HandleFunc("/list/", notImplemented)
	http.HandleFunc("/remove/", notImplemented)
	lg.Fatal(http.ListenAndServe(*addr, nil))
}

func fixPath(n int, fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = r.URL.Path[n:]
		fn(w, r)
	}
}

func notImplemented(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}
