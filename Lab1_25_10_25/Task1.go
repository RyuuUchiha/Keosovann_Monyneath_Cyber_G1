package main

import "fmt"

func task1() {
	fmt.Println("	Task 1: Assignment Operator")

	var num1, num2 int

	fmt.Print("Enter first integer: ")
	fmt.Scan(&num1)

	fmt.Print("Enter second integer: ")
	fmt.Scan(&num2)

	num3 := num1
	fmt.Printf("num3 = num1 : %d\n", num3)

	num3 += num2
	fmt.Printf("num3 += num2 : %d\n", num3)

	num3 -= num2
	fmt.Printf("num3 -= num2 : %d\n", num3)

	num3 *= num2
	fmt.Printf("num3 *= num2 : %d\n", num3)

	if num2 != 0{
		num3 /= num2
	fmt.Printf("num3 /= num2 : %d\n", num3)
	} else {
		fmt.Println("Error: division by zero")
	}
	
	if num2 != 0 {
	num3 %= num2
	fmt.Printf("num3 %%= num2 : %d\n", num3)
	} else 	{
		fmt.Println("Error: modulus by zero")
	}
}
