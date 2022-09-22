package main

import (
	"fmt"
	"math/rand"
	"time"
)

func displayHeading(){
	fmt.Println("Rock, Paper, Scissors!")
}

func getComputerChoice() string{
	var compChoice string
	min := 1
	max := 4

	rand.Seed(time.Now().UnixNano())
	randomInt := rand.Intn(max - min) + min

	if randomInt == 1 {
		compChoice = "rock"
	} else if randomInt == 2 {
		compChoice = "paper"
	} else {
		compChoice = "scissors"
	}
	return compChoice

}

func main(){
	displayHeading()
	fmt.Println("Hello, welcome to the game\n")

	fmt.Println("The game is about to begin....")

	var isRunning bool
	isRunning = true

	fmt.Println(getComputerChoice())

	for isRunning {
		fmt.Printf("Enter Rock, Paper or Scissor. Enter End to end the game: ")

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