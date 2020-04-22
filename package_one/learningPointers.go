package main

import "fmt"

func main() {
	x := 12 // 12 is stored in a memory location with x pointing to that location
	y := 22
	a := &x // a is now pointing to the memory location, x
	fmt.Printf("Memory location of x: %v\n", a)
	fmt.Printf("Value stored at memory x: %v\n", *a)
	*a = 200
	fmt.Printf("New value of x: %v\n", x)
	*a = *a * 2
	fmt.Printf("New value of x after multiplying by 2: %v\n", x)
	a = &y
	fmt.Printf("A is now pointing to y: %v\n", *a)
}
