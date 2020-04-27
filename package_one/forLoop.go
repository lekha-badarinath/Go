package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func traditionalForLoop() {
	arr := [5]int{1, 2, 3, 4, 5}
	var s string
	for i := 0; i < len(arr); i++ {
		s += strconv.Itoa(arr[i]) + " "
	}
	fmt.Println(s)
	fmt.Println(reflect.TypeOf(s))
}

func loopWithOnlyCondition() {
	n := 1
	for n <= 12 {
		fmt.Println(n)
		n += 1
	}
}

func infiniteWhileLoop() {
	n := 1
	fmt.Println("Coming from the infinite loop")
	for {
		fmt.Println(n)
		n++
		if n == 12 {
			break
		}
	}
}

func usingRangeAndBlankIdentifier() {
	arr := [5]string{"a", "b", "c", "d", "e"}
	s := ""
	/* the range keyword generated a index, value pair.
	Most of the times, we do not need the index value of a element.
	Since go does not allow use to leave local variables unused, 'a' would give rise to a compilation error
	To avoid this, we use the blank identifier _
	The blank identifier tells the Go compiler that there is some data here which we do not wish to use.*/
	for a, n := range arr[:len(arr)] { //or arr[0:]
		s += n + " "
		fmt.Println(a, n)
	}
	fmt.Println(s)
	// One line command to join the elements of the array.
	fmt.Println(strings.Join(arr[0:], " "))
	fmt.Println(arr[0:])
}

func main() {
	//infiniteWhileLoop()
	//traditionalForLoop()
	//loopWithOnlyCondition()
	usingRangeAndBlankIdentifier()
}
