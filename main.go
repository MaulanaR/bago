package main

import (
	"fmt"
	"log"
	"net/http"

	"rsc.io/quote"
)

func hello(w http.ResponseWriter, r *http.Request) {
	url := r.URL
	fmt.Fprintf(w, "hello from %v", url)
}
func awal(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HIDUP HIDUP HIDUP MAHASISWA")
}
func quotegolang(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, quote.Go())
}

func main() {
	ListenAddress := ":8090"
	http.HandleFunc("/", awal)
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/hello1", hello)
	http.HandleFunc("/quote", quotegolang)

	log.Println("Listen at localhost" + ListenAddress)
	http.ListenAndServe(ListenAddress, nil)
}
