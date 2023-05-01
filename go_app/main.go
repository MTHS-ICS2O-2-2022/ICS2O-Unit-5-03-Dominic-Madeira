// Created by: Dominic M.
// Created on: April 2023
//
// This program converts fahrenheit to celsius.
package main

import (
  "fmt"
)

func main() {
  var age int

  // input
  fmt.Println("This program finds your age rating for movies")
  fmt.Println()
  fmt.Print("Enter your age: ")
  fmt.Scanln(&age)

  // process
	if age >= 18 {
		fmt.Println()
		fmt.Println("You can go to an R rated movie alone")
	} else if age >= 13 {
		fmt.Println()
		fmt.Println("You can go to a PG-13 rated movie alone")
	} else {
		fmt.Println()
		fmt.Println("You can go to a G rated movie alone")
	}

  fmt.Println("\nDone.")
}
