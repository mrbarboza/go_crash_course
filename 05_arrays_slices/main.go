package main

import "fmt"

func main() {
	// Array
	// var fruitArr [2]string

	// Assign Values
	// fruitArr[0] = "Apple"
	// fruitArr[1] = "Orange"

	// Declare and Assign
	fruitArr := [2]string{"Apple", "Orange"}

	fmt.Println(fruitArr)
	fmt.Println(fruitArr[1])

	// Slice
	fruitSlice := []string{"Apple", "Orange", "Grape", "Cherry"}
	fmt.Println(fruitSlice)
	fmt.Println(fruitSlice[1:3])
}
