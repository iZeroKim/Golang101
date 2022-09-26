package main

import(
	"fmt"
)

type Movie struct{
	title string
	length string
	// To be changed to time type
	release_date string
	likes int
}


func main(){
	movie1 := Movie{"Last King", "2 hours", "26/09/2022"}	
	fmt.Println(movie1.title, movie1.length, movie1.release_date, movie1.likes)
}