package main

import (
	"bufio"
	"fmt"
	"os"
)

func task1() {

	fmt.Println("	📋Task 1: Assignment Operator")

	var num1, num2 int

	fmt.Print("Enter first integer: ")
	if _, err := fmt.Scan(&num1); err != nil {
		fmt.Println("❌ Invalid input! Please enter a number only.")
		bufio.NewReader(os.Stdin).ReadString('\n')
		return
	}

	fmt.Print("Enter second integer: ")
	if _, err := fmt.Scan(&num2); err != nil {
		fmt.Println("❌ Invalid input! Please enter a number only.")
		bufio.NewReader(os.Stdin).ReadString('\n')
		return
	}

	fmt.Println("\n--- Results ---")
	num3 := num1
	fmt.Printf("num3 = num1 : %d\n", num3)

	num3 += num2
	fmt.Printf("num3 += num2 : %d\n", num3)

	num3 -= num2
	fmt.Printf("num3 -= num2 : %d\n", num3)

	num3 *= num2
	fmt.Printf("num3 *= num2 : %d\n", num3)

	if num2 != 0 {
		num3 /= num2
		fmt.Printf("num3 /= num2 : %d\n", num3)
	} else {
		fmt.Println("⚠️ Error: division by zero‼️")
	}

	if num2 != 0 {
		num3 %= num2
		fmt.Printf("num3 %%= num2 : %d\n", num3)
	} else {
		fmt.Println("⚠️ Error: modulus by zero‼️")
	}
}
