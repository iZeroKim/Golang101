package main

import "fmt"

func displayHeading(){
	fmt.Println("Rock, Paper, Scissors!")
}
func main(){
	displayHeading()
	fmt.Println("Hello, welcome to the game\n")

	fmt.Println("The game is about to begin....")

	var isRunning bool
	isRunning = true

	for isRunning {
		fmt.Println("Enter Rock, Paper or Scissor. Enter End to end the game")

		var choice string
		fmt.Scan(&choice)

		if choice == "end" {
			fmt.Println("Game Ended")
			isRunning = false
		} else {
			fmt.Println(choice)
		}
	}

}