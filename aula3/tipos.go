package main

import "fmt"

func main() {
	// Bool
	fmt.Printf("Type:  %T - Value: %v\n ", true, true)
	fmt.Printf("Type:  %T - Value: %v\n ", false, false)

	// String
	fmt.Printf("Type:  %T - Value: %v\n ", "Will", "Will")
	fmt.Printf("Type:  %T - Value: %v\n ", "Will123", "Will123")
	fmt.Printf("Type:  %T - Value: %v\n ", "1", "1")
	
	// Int
	fmt.Printf("Type:  %T - Value: %v\n ", 1, 1)
	fmt.Printf("Type:  %T - Value: %v\n ", 12, 12)
	fmt.Printf("Type:  %T - Value: %v\n ", 125,125)
	// Float
	fmt.Printf("Type:  %T - Value: %v\n ", 1.233, 1.233)
	fmt.Printf("Type:  %T - Value: %v\n ", 12.5, 12.5)
}