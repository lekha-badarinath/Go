package main

import (
	"fmt"
	"math/rand"
)

func generateRandomNum() int {
	return rand.Intn(100)
}

func main() {
	fmt.Println("Random number: ", generateRandomNum())
}