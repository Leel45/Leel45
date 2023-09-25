package main

import (
	"fmt"
)

// Define the Animal interface
type Animal interface {
	Speak() string
}

// Define a Dog struct
type Dog struct{}

// Implement the Speak method for Dog
func (d Dog) Speak() string {
	return "Woof!"
}

// Define a Cat struct
type Cat struct{}

// Implement the Speak method for Cat
func (c Cat) Speak() string {
	return "Meow!"
}

func main() {
	// Create instances of Dog and Cat
	dog := Dog{}
	cat := Cat{}

	// Create a slice of Animal interfaces and add instances of Dog and Cat to it
	animals := []Animal{dog, cat}

	// Iterate over the slice of animals and let them speak
	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}
}
