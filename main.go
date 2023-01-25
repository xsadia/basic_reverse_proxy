package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	serverHandler := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		fmt.Printf("received request at %s", time.Now())
		fmt.Fprint(rw, "received with success\n")
	})

	log.Fatal(http.ListenAndServe(":8081", serverHandler))
}