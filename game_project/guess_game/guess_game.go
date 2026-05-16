package guess_game

import (
	"fmt"
	"math/rand"
	"time"
)

func StartGame() {
	random_number := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(100) + 1
	var guess int
	fmt.Print("Welcome to the Guess Game! Try to guess the number between 1 and 100: ")
	fmt.Scanf("%d", &guess)
	for guess != random_number && guess != 0 {
		if guess == 0 {
			fmt.Println("Game exited. The number was:", random_number)
			return
		} else {
			if guess < random_number {
				fmt.Println("Too low! Try again: ")
			} else {
				fmt.Println("Too high! Try again: ")
			}
			fmt.Scanf("%d", &guess)
		}
	}
	fmt.Println("Congratulations! You guessed the number!", random_number)
}
