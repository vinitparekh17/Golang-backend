package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// http.HandleFunc("/", serveHome)
	// fmt.Println("Server is running on port 8080")

	// now doing the same thing using goroutines which is used to run multiple functions at the same time (lightweight threads)
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome)
	log.Fatal(http.ListenAndServe(":8080", r))
	fmt.Println("Server is running on port 8080")
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Triggered serveHome")
	w.Write([]byte("<font color='#00ffab'><h1>Welcome to home page</h1></font>"))
}
