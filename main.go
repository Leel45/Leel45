package main

import (
	"fmt"
)

// Define the Company struct
type Company struct {
	Name        string
	Location    string
	Departments []string
}

// Method to add a department to the Company's Departments slice
func (c *Company) AddDepartment(department string) {
	c.Departments = append(c.Departments, department)
}

// Function to trim a department from the Company's Departments slice
func (c *Company) TrimDepartment(departmentToRemove string) {
	newDepartments := []string{}
	for _, department := range c.Departments {
		if department != departmentToRemove {
			newDepartments = append(newDepartments, department)
		}
	}
	c.Departments = newDepartments
}

func main() {
	// Create an instance of the Company struct
	company := Company{
		Name:     "Example Inc.",
		Location: "New York",
	}

	// Add departments to the company
	company.AddDepartment("HR")
	company.AddDepartment("Finance")
	company.AddDepartment("Engineering")

	// Print the company's information
	fmt.Printf("Company Name: %s\n", company.Name)
	fmt.Printf("Location: %s\n", company.Location)
	fmt.Printf("Departments: %v\n", company.Departments)

	// Trim a department from the company
	company.TrimDepartment("HR")

	// Print the updated departments
	fmt.Printf("Updated Departments: %v\n", company.Departments)
}
