package main

import (
	"fmt"
)

// Original struct
type Person struct {
	Name     string
	Age      int
	Location string
}

// Extended struct
type Employee struct {
	Person         // Embedding Person struct
	EmployeeID int // Additional field for Employee
	Role       string
}

func main() {
	// Creating a new Person instance
	brian := Person{
		Name:     "Brian Su",
		Age:      40,
		Location: "Toronto",
	}

	// Accessing fields from Person
	fmt.Println("Name:", brian.Name)
	fmt.Println("Age:", brian.Age)
	fmt.Println("Location:", brian.Location)

	fmt.Println("------")

	// Creating a new Employee instance
	emp := Employee{
		Person: Person{
			Name:     "John Doe",
			Age:      30,
			Location: "New York",
		},
		EmployeeID: 1001,
		Role:       "Software Engineer",
	}

	// Accessing fields from both Person and Employee
	fmt.Println("Name:", emp.Name)
	fmt.Println("Age:", emp.Age)
	fmt.Println("Location:", emp.Location)
	fmt.Println("EmployeeID:", emp.EmployeeID)
	fmt.Println("Role:", emp.Role)
}
