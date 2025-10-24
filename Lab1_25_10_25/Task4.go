package main

import (
	"fmt"
)

func task4() {
	fmt.Println("	📋Task4: Mini Calculator")
	home()
}

func home() {
	fmt.Println("==============================")
	fmt.Println("===== Mini Calculator 🧮 =====")
	fmt.Println("==============================")
	fmt.Println("1) Addition ➕")
	fmt.Println("2) Substraction ➖")
	fmt.Println("3) Multiplication ✖️")
	fmt.Println("4) Division ➗")
	fmt.Println("5) Modulus 🧩")
	fmt.Println("6) Exit 🚪")
	fmt.Print("Please choose the operation: ")
	var choice int
	fmt.Scan(&choice)
	if choice == 6 {
		fmt.Println("Goodbye!👋 Good luck🍀")
		return
	}
	var c, d float64
	fmt.Print("Enter first number: ")
	fmt.Scan(&c)
	fmt.Print("Enter second number: ")
	fmt.Scan(&d)
	switch choice {
	case 1:
		fmt.Printf("Result: %.2f + %.2f = %.2f\n", c, d, add(c, d))
		home()
	case 2:
		fmt.Printf("Result: %.2f - %.2f = %.2f\n", c, d, sub(c, d))
		home()
	case 3:
		fmt.Printf("Result: %.2f * %.2f = %.2f\n", c, d, mul(c, d))
		home()
	case 4:
		div(c, d)
		home()
	case 5:
		mod(c, d)
		home()
	default:
		fmt.Println("Invalid choice, try again. 🔄️")
	}
}

func add(c, d float64) float64 {
	return c + d
}

func sub(c, d float64) float64 {
	return c - d
}

func mul(c, d float64) float64 {
	return c * d
}

func div(c, d float64) {
	if d == 0 {
		fmt.Println("Error: division by zero‼️")
		return
	}
	fmt.Printf("Result: %.2f / %.2f = %.2f\n", c, d, c/d)
}

func mod(c, d float64) {
	if d == 0 {
		fmt.Println("Error: modulus by zero‼️")
		return
	}
	fmt.Printf("Result: %.2f %% %.2f = %.2f\n", c, d, float64(int(c)%int(d)))
}
