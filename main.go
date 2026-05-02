package main

import (
	"fmt"
	"log"
	"net/http"
)

//the helloHandler handles the GET request from the /hello path
func helloHandler(w http.ResponseWriter, r *http.Request) {
 //verify if the path is equal to /hello
  if r.URL.Path != "/hello" {
    http.Error(w, "404 not found", http.StatusNotFound)
    return 
  }

  //verify if the method is equal to GET
  if r.Method != "GET" {
    http.Error(w, "Method is not supported", http.StatusMethodNotAllowed)
    return
  }

  fmt.Fprint(w, "hello!") 
}

//the formHandler handles the POST request from the form
func formHandler(w http.ResponseWriter, r *http.Request) {
  if err := r.ParseForm(); err != nil {
    fmt.Fprint(w, "Error parsing form: %v\n", err)
    return
  }

  fmt.Fprintf(w, "POST request received!")

  name := r.FormValue("name")
  address := r.FormValue("address")

  fmt.Fprintf(w, "Name: %s\n", name)
  fmt.Fprintf(w, "Address: %s\n", address)

}

func main() {
	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileserver)

	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("\nStarting the server on the Port 8080\n")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
