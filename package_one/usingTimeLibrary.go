package main

import (
	"fmt"
	"strings"
	"time"
)

func usingForLoop() float64 {
	s := [5]string{"a", "b", "c", "d", "e"}
	str := ""
	start := time.Now()
	for _, i := range s[0:] {
		str += i + " "
	}
	fmt.Println(str)
	stop := time.Since(start).Seconds()
	return stop
}

func usingStringLibrary() float64 {
	s := [5]string{"a", "b", "c", "d", "e"}
	start := time.Now()
	fmt.Println(strings.Join(s[0:], " "))
	stop := time.Since(start).Seconds()
	return stop
}

func main() {
	fmt.Printf("Time taken using for loop:%f\n", usingForLoop())
	fmt.Printf("Time taken using \"strings.Join\" function:%f\n", usingStringLibrary())
}
