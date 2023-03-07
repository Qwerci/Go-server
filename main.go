package main

import(
	"fmt"
	"log"
	"net/http"
)



func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path!= "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "405 method not allowed",http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Hello, World!\n")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm error: %v\n", err)
		return
	}

	fmt.Fprint(w, "POST request received!\n")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "name: %s\n", name)
	fmt.Fprintf(w, "address: %s\n", address)
}

func main() {
	fileserver := http.FileServer(http.Dir("/.static-folder"))
	http.Handle("/", fileserver)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("starting server on :8080\n")
	if err := http.ListenAndServe(":8080", nil); err!= nil { 
		log.Fatal(err)
	}
}