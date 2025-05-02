package main

import (
	"fmt"
	"log"
	"net/http"
)


func apiHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}

func startFrontendServer() {
	fs := http.FileServer(http.Dir("frontend"))
	http.Handle("/", fs)
	log.Fatal(http.ListenAndServe(":80", nil))
}

func startBackendServer() {
	http.HandleFunc("/api", apiHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func main() {
	go startFrontendServer()
	startBackendServer()
}