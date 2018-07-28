package main

import (
	"fmt"
	"log"
	"net/http"
)

func SayHello(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "Host/port: %s\n", request.Host)
	fmt.Fprintf(response, "Method: %s\n", request.Method)
	fmt.Fprintf(response, "Proto: %s\n", request.Proto)
	fmt.Fprintf(response, "URI: %s\n", request.RequestURI)
	for _, v := range request.Cookies() {
		fmt.Fprintf(response, "Cookies: %s\n", v)
	}

	response.Write([]byte("Hello World!"))
}

func main() {
	http.HandleFunc("/hello", SayHello)
	log.Fatal(http.ListenAndServe(":8010", nil))
}
