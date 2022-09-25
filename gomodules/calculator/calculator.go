package main

import (
	"fmt"
	"example.com/addition"
)

func main(){
	fmt.Println("\n\nCalculator App - Addition only")
	
	var firstNum, secondNum int
	fmt.Printf("First number: ")
	fmt.Scan(&firstNum)

	fmt.Printf("Second number: ")
	fmt.Scan(&secondNum)

	fmt.Printf("Sum is : %v", addition.Add(firstNum, secondNum))

}
