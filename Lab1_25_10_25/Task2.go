package main

import (
	"bufio"
	"fmt"
	"os"
)

func task2() {

	fmt.Println("	📋Task 2: Logical Operator")
	var num4, num5 int

	fmt.Print("Enter first integer: ")
	if _, err := fmt.Scan(&num4); err != nil {
		fmt.Println("❌ Invalid input! Please enter a number only.")
		bufio.NewReader(os.Stdin).ReadString('\n')
		return
	}

	fmt.Print("Enter second integer: ")
	if _, err := fmt.Scan(&num5); err != nil {
		fmt.Println("❌ Invalid input! Please enter a number only.")
		bufio.NewReader(os.Stdin).ReadString('\n')
		return
	}

	fmt.Println("\n--- Results ---")
	and := num4 > 0 && num5 > 0
	fmt.Printf("%d > 0 && %d > 0 → %t\n", num4, num5, and)

	or := num4 > num5 || num4 < num5
	fmt.Printf("%d > %d || %d < %d → %t\n", num4, num5, num4, num5, or)

	not := num4 != num5
	fmt.Printf("%d != %d → %t\n", num4, num5, not)
}
