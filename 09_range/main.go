package main

import "fmt"

func main() {
	ids := []int{33, 76, 54, 23, 11, 2, 1}

	// Loop through ids
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// Not using index
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	// Add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum", sum)

	// Range with map
	emails := map[string]string{"Miguel": "miguel@email.com", "Akemi": "akemi@email.com"}

	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}

	for k := range emails {
		fmt.Printf("Name: %s\n", k)
	}

	for _, v := range emails {
		fmt.Printf("Email: %s\n", v)
	}
}
