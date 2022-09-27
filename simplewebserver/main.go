package main

import (
	"fmt"
	"net/http"
)

func main(){
	fmt.Println("Simple Go Web Server")
	http.ListenAndServe(":3000", nil)
}