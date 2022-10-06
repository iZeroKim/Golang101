package main

import ("net/http"
		"fmt"
)

type Info struct{
	Title string
	Name string
	Fruits [3]string
}

func main(){
	fmt.Println("Starting new server at: http://localhost:8000")
	http.ListenAndServe(":8000", nil)
}