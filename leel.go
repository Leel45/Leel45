package main

import "fmt"

func sumAndCount(numbers ...int) (int, int) {
	sum := 0
	count := len(numbers)
	for _, num := range numbers {
		sum += num
	}
	return sum, count
}

func main() {
	result, count := sumAndCount(1, 2, 3, 4, 5)
	fmt.Printf("Sum: %d, Count: %d\n", result, count)
}
