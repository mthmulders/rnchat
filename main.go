package main

import (
	"flag"
	"log"
)

var addr = flag.String("addr", ":8080", "service address")

func main() {
	flag.Parse()
	log.Printf("Listening at %v", *addr)
}
