package main

import (
	"fmt"
)

func shortHandType() string {
	statement := "Hey there! Welcome to Go"
	return statement
}

func main() {
	statement := shortHandType()
	fmt.Println(statement)
	fmt.Printf("Type of statement is: %T", statement )
}
