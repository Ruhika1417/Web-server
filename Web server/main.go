package main

import (
	"fmt"
	"log"
	"net/http"
)

// to a server, a request sent by user and a response sent by server.

func FormHandler(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "wrong details: %v", err)
	}

	// fmt.Println("post request successful")
	fmt.Fprint(w, "post request succesful\n")
	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address= %s\n", address)

}

func HelloHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/hello" {
		http.Error(w, "error 404 not found", http.StatusBadRequest)
	}

	if r.Method != "GET" {
		http.Error(w, "bad request", http.StatusBadRequest)
	}

	// fmt.Println("Hello")
	fmt.Fprintf(w, "hellu!!")

}

func main() {

	FileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", FileServer)
	http.HandleFunc("/hello", HelloHandler)
	http.HandleFunc("/form", FormHandler)

	fmt.Println("Server starting at port 8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

}
