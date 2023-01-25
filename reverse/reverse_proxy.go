package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

func main() {
	originServerUrl, err := url.Parse("http://127.0.0.1:8081")

	if err != nil {
		log.Fatal("invalid origin server url")
	}

	reverseProxy := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		fmt.Printf("received request at: %s", time.Now())

		req.Host = originServerUrl.Host
		req.URL.Host = originServerUrl.Host
		req.URL.Scheme = originServerUrl.Scheme
		req.RequestURI = ""

		originServerResponse, err := http.DefaultClient.Do(req)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(rw, err)
			return 
		}

		rw.WriteHeader(http.StatusOK)
		io.Copy(rw, originServerResponse.Body)
	})

	log.Fatal(http.ListenAndServe(":8080",reverseProxy))
}