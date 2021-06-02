package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

func (person Person) greet() string {
	return "Hello, my name is " + person.firstName + " " + person.lastName + " and I am " + strconv.Itoa(person.age)
}

func (person *Person) hasBirthday() {
	person.age++
}

func (person *Person) getMaried(spouseLastName string) {
	if person.gender == "M" {
		return
	} else {
		person.lastName = spouseLastName
	}
}

func main() {
	person1 := Person{
		firstName: "Akemi",
		lastName:  "Medeiros",
		city:      "Santo André",
		gender:    "F",
		age:       34,
	}

	fmt.Println(person1)
	fmt.Println(person1.firstName)
	person1.age++
	fmt.Println(person1)

	fmt.Println(person1.greet())
	person1.hasBirthday()
	person1.getMaried("Barboza")
	fmt.Println(person1.greet())
}
