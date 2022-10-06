package main

import ("net/http"
		"fmt"
)

func main(){
	fmt.Println("Starting new server at: http://localhost:8000")
	http.ListenAndServe(":8000", nil)
}