package main

import "fmt"

type person struct {
	name       string
	age        uint
	profession string
}

// Method returning a value - Value pointers
func (per person) helloMessage() (string, uint) {
	return per.name, per.age
}

func main() {
	personA := person{
		name:       "Chris",
		age:        38,
		profession: "Actor",
	}
	fmt.Println(personA.helloMessage())
}
