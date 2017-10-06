package main

import (
    "fmt"
    "net/http"
    "os"
)

func shippingHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Shipping microservice version", os.Getenv("VERSION"))
}

func main() {
	fmt.Println("Starting web-shipping version", os.Getenv("VERSION"))
    http.HandleFunc("/", shippingHandler)
    http.ListenAndServe(":8080", nil)
}
