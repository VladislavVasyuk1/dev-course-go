package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello1, 世界")
}

func handlerNames(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Vlada, Vlad")
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/names", handlerNames)
	fmt.Println("Running demo app. Press Ctrl+C to exit...")
	log.Fatal(http.ListenAndServe(":8888", nil))
}
