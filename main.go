package test_sonar_go

import (
	"fmt"
	"math/rand"
)

func main() {
	randomNumber := rand.Intn(100) // Generates a random number between 0 and 99
	fmt.Println("Random Number:", randomNumber)
}
