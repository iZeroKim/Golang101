package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "<h1>Hello World</h1>")
}

func main(){
	fmt.Println("Simple Go Web Server")
	http.HandleFunc("/", index)
	fmt.Println("Server starting...")
	http.ListenAndServe(":3000", nil)
	fmt.Println("Server running...")
}

