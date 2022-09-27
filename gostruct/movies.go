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

//Constructor function
func(movie *Movie) setDefaults(){
	if movie.likes == 0{
		movie.likes = 0
	}
}

//value receiver struct function
func(movie Movie) displayMovie(){
	fmt.Printf("Name: %v, Duration: %v, Date Released: %v, Likes: %v", movie.title, movie.length, movie.release_date, movie.likes)
}


func main(){
	movie1 := Movie{title:"Last King", length:"2 hours", release_date:"26/09/2022"}	
	movie1.setDefaults()
	movie1.displayMovie()
}