package main

import "fmt"

func main() {
	// Define map
	// emails := make(map[string]string)

	// Assign KV
	// emails["Miguel"] = "miguel@email.com"
	// emails["Akemi"] = "akemi@email.com"
	// emails["Hiroshi"] = "hiroshi@email.com"

	// Declare a map and assign KV
	emails := map[string]string{"Miguel": "miguel@email.com", "Akemi": "akemi@email.com"}
	emails["Hiroshi"] = "hiroshi@email.com"

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["Akemi"])

	// Delete from map
	delete(emails, "Akemi")
	fmt.Println(emails)
}
