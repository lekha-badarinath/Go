 package main

import (
	"../package_A"
	"../package_B"
	"fmt"
)
// Several packages can be called in one "main" package
// There can only be one main package
func main() {
	var quo float64 = package_A.DivideTwoNumbers(12,8)
	var prod float64 = package_B.MultiplyTwoNumbers(16,234)
	var sum int = package_A.AddTwoNumbers(224, 86)
	fmt.Printf("Divide two numbers: %f\nMultiply two numbers: %f\nAdd two numbers: %v\n", quo, prod, sum)
}
