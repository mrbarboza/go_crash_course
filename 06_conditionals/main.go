package main

import "fmt"

func main() {
	x := 10
	y := 10

	// If Else
	fmt.Println("If/Else Statement.")
	if x <= y {
		fmt.Printf("%d is less or equal to %d.\n", x, y)
	} else {
		fmt.Printf("%d is less than %d.\n", y, x)
	}

	// Else If
	fmt.Println("Else/If Statement.")
	color := "green"

	if color == "red" {
		fmt.Println("the color is red.")
	} else if color == "blue" {
		fmt.Println("the color is red.")
	} else {
		fmt.Println("the color is NOT blue or red.")
	}

	// Switch
	fmt.Println("Switch Statement.")
	switch color {
	case "red":
		fmt.Println("the color is red.")
	case "blue":
		fmt.Println("the color is blue.")
	default:
		fmt.Println("the color is NOT blue or red.")
	}
}
