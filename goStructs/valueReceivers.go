package main

import "fmt"

type employee struct {
	employeeName string
	role         string
	salary       float64
}

const hikePercentage float64 = 12

// method - value receivers
func (emp employee) calculateHike() float64 {
	return emp.salary * hikePercentage / 100
}

func (emp employee) newSalary() float64 {
	return emp.salary + emp.calculateHike()
}

func main() {
	employeeA := employee{
		employeeName: "Chris",
		role:         "Developer",
		salary:       150000,
	}
	fmt.Println(employeeA.calculateHike())
	fmt.Println(employeeA.newSalary())
}
