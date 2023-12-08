package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	n, err := fmt.Fprintf(w, "This is the home page")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(fmt.Sprintf("No of bytes written:%d", n))
}

func About(w http.ResponseWriter, r *http.Request) {
	n, err := fmt.Fprintf(w, "This is the about page")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(fmt.Sprintf("No of bytes written:%d", n))
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Starting Server on port %s", portNumber))

	http.ListenAndServe(portNumber, nil)
}
