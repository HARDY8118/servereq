package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
)

const VERSION = "0.0.1"

type requestLogger struct{}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func handler(w http.ResponseWriter, req *http.Request) {
	// Request method and URL
	fmt.Printf("%s %s %s%s\n", req.Proto, req.Method, req.Host, req.URL.Path)

	// Headers
	for name, headers := range req.Header {
		fmt.Printf("%s: %s\n", name, headers)
	}

	// Body
	body := new(bytes.Buffer)
	body.ReadFrom(req.Body)
	fmt.Println(body.String())
	fmt.Println()
}

func main() {

	http.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Version: %s", VERSION)
	})

	http.HandleFunc("/", handler)

	if len(os.Args) < 2 {
		fmt.Println("Need port number to start server")
		os.Exit(1)
	} else {
		PORT := os.Args[1]
		fmt.Printf("Server starting on PORT: %s\n\n", PORT)
		http.ListenAndServe(":"+PORT, nil)
	}
}
