package main

import "fmt"

func main() {
	// If-else-if condition
	temperature := 75

	if temperature > 90 {
		fmt.Println("It's hot outside!")
	} else if temperature > 70 {
		fmt.Println("It's warm outside.")
	} else {
		fmt.Println("It's cool outside.")
	}

	// Switch statement with multiple switches
	day := "Wednesday"
	switch day {
	case "Monday", "Tuesday", "Wednesday":
		fmt.Println("It's a weekday.")
	case "Thursday", "Friday":
		fmt.Println("It's almost the weekend!")
	case "Saturday", "Sunday":
		fmt.Println("It's the weekend!")
	default:
		fmt.Println("Invalid day.")
	}

	// For loop without initialization statement, using break and continue
	i := 0
	for {
		if i >= 5 {
			break
		}
		if i%2 == 0 {
			i++
			continue
		}
		fmt.Println(i)
		i++
	}

}
