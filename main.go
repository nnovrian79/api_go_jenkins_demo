package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World apa kabar nova")
	fmt.Fprintf(w, "Baris ini ditulis oleh nova")
	fmt.Fprintf(w, "Hello, World apa kabar sumar")
	fmt.Fprintf(w, "Baris ini ditulis oleh sumar")
	fmt.Fprintf(w, "Hello, World apa kabar raki")
	fmt.Fprintf(w, "Baris ini ditulis oleh raki")
	

}

func main() {
	port := ":8070"

	// Print a message indicating the server is starting
	fmt.Printf("Starting server on port%s\n", port)
	http.HandleFunc("/", handler)

	http.ListenAndServe(port, nil)
	fmt.Printf("Appication stop\n")
}
