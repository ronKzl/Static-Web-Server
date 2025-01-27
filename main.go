package main

import (
	"fmt"
	"log"
	"net/http"
)

// helloHandler handles route requests to /hello and returns a string to the user if access is of the correct type and succesfull
// res - the response field that will be returned
// req - the request field ptr that stores user recieved info
func helloHandler(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/hello" {
		http.Error(res, "404 Not Found", http.StatusNotFound)
		return
	}

	if req.Method != "GET" {
		http.Error(res, "Method is not supported", http.StatusNotFound)
		return
	}

	fmt.Fprintf(res, "Hello")
}

func formHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	//serve index.html base route
	http.Handle("/", fileServer)
	//serve and process the form.html
	http.HandleFunc("/form", formHandler)
	//do hello func
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server on port 8080 \n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
