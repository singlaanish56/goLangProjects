package main

import (
	"fmt"
	"strings"
	"log"
	"net/http"
)

func myHandler(w http.ResponseWriter, r * http.Request){
	response := []string{
		"Received request from " + r.RemoteAddr,
		"GET / HTTP/1.1",
		"Host: localhost",
		"User-Agent: curl/7.85.0",
		"Accept: */*",
	}

	fmt.Println(strings.Join(response, "\r\n"))
}
func main() {
	fmt.Println("Welcome to the loadBalancer")

	http.HandleFunc("/", myHandler)

	log.Fatal(http.ListenAndServe(":80", nil))
}