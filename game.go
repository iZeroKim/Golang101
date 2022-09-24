package main

import (
	"fmt"
	"math/rand"
	"time"
	"strings"
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
	} else if randomInt == 3 {
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

	var choice string
	for isRunning {
		fmt.Printf("Enter Rock, Paper or Scissors. Enter End to end the game: ")

		
		fmt.Scan(&choice)
		choiceSmall := strings.ToLower(choice)
		computerChoice := getComputerChoice()
		if choiceSmall == "end" {
			fmt.Println("Game Ended")
			isRunning = false
		} else {
			if choiceSmall == "rock" || choiceSmall == "paper" || choiceSmall ==  "scissors"{
				fmt.Println("You chose : ", choice)
				fmt.Println("The computer chose : ", computerChoice)
				if choiceSmall == computerChoice{
					fmt.Printf("You both selected %v. It's a tie!\n", choiceSmall)
				} else if choiceSmall == "rock"{
					if computerChoice == "scissors"{
						fmt.Println("Rock smashes scissors! You win!")
					} else{
						fmt.Println("Paper covers rock! You lose.")
					}
				} else if choiceSmall == "paper"{
					if computerChoice == "rock"{
						fmt.Println("Paper covers rock! You win!")
					}else{
						fmt.Println("Scissors cuts paper! You lose.")
					}
				} else if choiceSmall == "scissors"{
					if computerChoice == "paper"{
						fmt.Println("Scissors cuts paper! You win!")
					}else{
						fmt.Println("Rock smashes scissors! You lose.")
					}
				}	
				fmt.Println()


			} else {
				fmt.Println("Invalid option. Please select again. ")
			}
			
		}
	}

}