package main

import "fmt"

func testFunction(x, y int) int {
	fmt.Println("Function in the same file")
	return x + y
}
func main() {
	// functions within the same package can be called without importing as they are all visible to one another
	fmt.Println(newFunction())
	fmt.Println(testFunction(12, 15))

}
