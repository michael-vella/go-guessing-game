package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
)

const lives = 3

func main() {
	var g int
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Generating a random number between 0 and 10...")

	n := rand.IntN(10)

	fmt.Printf("You have %d lives to guess the random number generated...\n", lives)

	for i := 0; i < lives; i++ {
		cl := i + 1

		fmt.Printf("\nYou are on life %d out of %d...\n", cl, lives)
		fmt.Print("Enter your guess: ")

		_, err := fmt.Scanln(&g)

		if err != nil {
			fmt.Println("Invalid input! No lives will be lost but please enter a valid integer between 0-9...")

			// clear input buffer
			reader.ReadString('\n')

			i--
			continue
		}

		if g < 0 || g >= 10 {
			fmt.Println("Out of bounds! No lives will be lost but please enter a valid integer between 0-9...")

			i--
			continue
		}

		if n == g {
			fmt.Printf("Correct guess. You win! You guessed %d and the random generated number was %d!\n", g, n)
			break
		} else {
			fmt.Println("Incorrect guess...")
		}

		if lives == cl {
			fmt.Println("WOMP WOMP. GAME OVER! Better luck next time.")
		}
	}

	fmt.Println("\nThank you for playing! Hope to see you again soon...")
	os.Exit(0)
}
