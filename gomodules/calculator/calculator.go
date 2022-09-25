package main

import (
	"fmt"
	"example.com/addition"
)

func main(){
	fmt.Println("Calculator App - Addition only")
	
	var firstNum, secondNum int
	fmt.Printf("First number: ")
	fmt.Scan(&firstNum)

	fmt.Printf("Second number: ")
	fmt.Scan(&secondNum)

}
