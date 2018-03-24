package main

import "sync"

var fsdb sync.Map

func fsload(k string) ([]byte, bool) {
	if v, ok := fsdb.Load(k); ok {
		return v.([]byte), ok
	}
	return nil, false
}

func fsstore(k string, v []byte) {
	fsdb.Store(k, v)
}
