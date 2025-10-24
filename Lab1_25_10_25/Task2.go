package main

import "fmt"

func task2(){
	fmt.Println("	Task 2: Logical Operator")
	var num4, num5 int

	fmt.Print("Enter first integer: ")
	fmt.Scan(&num4)

	fmt.Print("Enter second integer: ")
	fmt.Scan(&num5)

	// Logical AND
	and := num4 > 0 && num5 > 0
	fmt.Printf("%d > 0 && %d > 0 → %t\n", num4, num5, and)

	// Logical OR
	or := num4 > num5 || num4 < num5
	fmt.Printf("%d > %d || %d < %d → %t\n", num4, num5, num4, num5, or)

	// Logical NOT (not equal)
	not := num4 != num5
	fmt.Printf("%d != %d → %t\n", num4, num5, not)
}