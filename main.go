package main

import (
	"fmt"
	"math/rand/v2"
	"os"
)

const lives = 1

func main() {
	var g int

	fmt.Println("Generating a random number between 0 and 10...")

	n := rand.IntN(10)

	fmt.Printf("You have %d life to guess the random number generated...\n", lives)

	fmt.Print("Enter your guess: ")
	_, err := fmt.Scan(&g)

	if err != nil {
		fmt.Println("Invalid input! Please enter a valid integer.")
		os.Exit(1)
	}

	if n == g {
		fmt.Printf("Correct guess! You guessed %d and the random generated number was %d!", g, n)
	} else {
		fmt.Printf("Incorrect guess. You guessed %d but the generated random number was %d...", g, n)
	}
}
