package main

import (
	"fmt"
)

func main(){
	var remainingTicket = 20 // Implicit Coversion of variable type 
	var totalTicket int = 100 // Explicit Declartion of variable type
	userName := "Mayur" // Implicit Conversion of Variable type
	fmt.Printf("Welcome %v",userName)
	fmt.Printf("Remaining ticket for the show %v and total ticket are %v",remainingTicket,totalTicket)
}