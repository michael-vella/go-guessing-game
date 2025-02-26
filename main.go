package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"slices"
)

func main() {
	var lives int
	var guess int
	var previousGuesses []int
	var upperBound int

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome! Let's play a quick game...")

	for {
		fmt.Print("\nPick any random number greater than 1: ")
		_, err := fmt.Scan(&upperBound)

		if err != nil {
			fmt.Println("Invalid input! Please enter a valid number...")
			reader.ReadString('\n') // clear input buffer
			continue
		}

		if upperBound <= 1 {
			fmt.Println("Invalid input! Number should be greater than 1...")
			continue
		}

		break
	}

	reader.ReadString('\n') // clear input buffer

	fmt.Printf("Generating a random number between 0 and %d (%d is included)...\n", upperBound, upperBound)
	n := rand.IntN(upperBound + 1)

	lives = GetLives(upperBound)
	fmt.Printf("You have %d lives to guess the random number generated...\n", lives)

	for i := 0; i < lives; i++ {
		cl := i + 1

		fmt.Printf("\nYou are on life %d out of %d...\n", cl, lives)
		fmt.Print("Enter your guess: ")

		_, err := fmt.Scanln(&guess)

		if err != nil {
			fmt.Printf("Invalid input! No lives will be lost but please enter a valid integer between 0-%d...\n", upperBound)
			reader.ReadString('\n') // clear input buffer
			i--
			continue
		}

		if guess < 0 || guess > upperBound {
			fmt.Printf("Out of bounds! No lives will be lost but please enter a valid integer between 0-%d...\n", upperBound)
			i--
			continue
		}

		if slices.Contains(previousGuesses, guess) {
			fmt.Printf("The number '%d' has already been guessed. Previous guesses: '%d'. Please enter a different number...\n", guess, previousGuesses)
			i--
			continue
		}

		previousGuesses = append(previousGuesses, guess)

		if guess == n {
			fmt.Printf("Correct guess. You win! You guessed %d and the random generated number was %d!\n", guess, n)
			break
		}

		if lives == cl {
			fmt.Printf("WOMP WOMP. GAME OVER! Better luck next time. The number was '%d'.", n)
			break
		}

		if guess < n {
			fmt.Println("Go higher!")
		} else {
			fmt.Println("Go lower!")
		}
	}

	fmt.Println("\nThank you for playing! Hope to see you again soon...")
	os.Exit(0)
}

func GetLives(n int) int {
	if n >= 1 && n <= 2 {
		return 1
	}
	return 1 + GetLives((n+1)/2) // ceiling division
}
