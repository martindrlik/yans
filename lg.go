package main

import (
	"log"
	"os"
)

var lg = log.New(os.Stdout, "yans: ", log.Ldate|log.Ltime|log.Lshortfile)
