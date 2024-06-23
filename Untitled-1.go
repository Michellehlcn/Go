package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
)

import "rsc.io/quote"

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, playground")
		fmt.Println(quote.Go())
	})

	log.Println("Starting server...")
	l, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		log.Fatal(http.Serve(l, nil))
	}()

	log.Println("Sending request...")
	res, err := http.Get("http://localhost:8080/hello")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Reading response...")
	if _, err := io.Copy(os.Stdout, res.Body); err != nil {
		log.Fatal(err)
	}
}
