package main

import "fmt"

type employee struct {
	employeeName string
	role         string
	salary       float64
}

const hikePercent = 12

// Value receiver
func (emp employee) calcuatedHike() float64 {
	return emp.salary * hikePercent / 100
}

// Pointer receiver
func (emp *employee) updateSalary(updatedSalary float64) {
	emp.salary = updatedSalary
}

/* Value receivers can be converted to pointer receivers by adding * before
the struct name.
Value receivers create a copy of the struct, which in turn takes up more memory*/

func main() {
	empA := employee{
		employeeName: "Chris",
		role:         "Developer",
		salary:       150000,
	}
	fmt.Println(empA.calcuatedHike())
	empA.updateSalary(160000)
	fmt.Println(empA)
	fmt.Println(empA.calcuatedHike())
}
