package main

import (
	"fmt"
)

func PrintType(x interface{}) {
	switch x := x.(type) {
	case int:
		fmt.Printf("x is an integer: %d\n", x)
	case float64:
		fmt.Printf("x is a float64: %f\n", x)
	case string:
		fmt.Printf("x is a string: %s\n", x)
	default:
		fmt.Println("Unknown type")
	}
}

func main() {
	var a = 42
	var b = 3.14
	var c = "Hello, World!"
	var d = true

	PrintType(a)
	PrintType(b)
	PrintType(c)
	PrintType(d)
}
