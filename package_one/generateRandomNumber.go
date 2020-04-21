package main

import (
	"math/rand"
)

func generateRandomNum() int {
	return rand.Intn(100)
}

//func main() {
//	num := generateRandomNum()
//	fmt.Println("Random number: ", num )
//}