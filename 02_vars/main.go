package main

import "fmt"

func main() {
	// MAIN Types
	// string
	// bool
	// int
	// int int8 int16 int32 int64
	// uint uint8 uint16 uint32 uint64
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128

	// Using var
	// var name = "Miguel"
	var age = 33
	var size float32 = 2.3

	// Shorthand
	name, email := "Miguel", "email@gmail.com"

	// const
	const isCool = true

	fmt.Println(name, age, isCool, email)
	fmt.Printf("%T\n", size)

}
