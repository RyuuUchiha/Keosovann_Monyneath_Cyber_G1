package main

import (
	"bufio"
	"fmt"
	"os"
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

	var choice int
	fmt.Print("Please choose the operation: ")
	if _, err := fmt.Scan(&choice); err != nil {
		fmt.Println("❌ Invalid input! Please enter a number between 1–6.")
		bufio.NewReader(os.Stdin).ReadString('\n')
		home()
		return
	}

	if choice < 1 || choice > 6 {
		fmt.Println("⚠️ Invalid choice! Please enter a number between 1–6.")
		home()
		return
	}

	if choice == 6 {
		fmt.Println("Goodbye!👋 Good luck🍀")
		return
	}

	var c, d float64
	fmt.Print("Enter first number: ")
	if _, err := fmt.Scan(&c); err != nil {
		fmt.Println("❌ Invalid input! Please enter a valid number.")
		bufio.NewReader(os.Stdin).ReadString('\n')
		home()
		return
	}

	fmt.Print("Enter second number: ")
	if _, err := fmt.Scan(&d); err != nil {
		fmt.Println("❌ Invalid input! Please enter a valid number.")
		bufio.NewReader(os.Stdin).ReadString('\n')
		home()
		return
	}

	switch choice {
	case 1:
		fmt.Printf("Result: %.2f + %.2f = %.2f\n", c, d, add(c, d))
	case 2:
		fmt.Printf("Result: %.2f - %.2f = %.2f\n", c, d, sub(c, d))
	case 3:
		fmt.Printf("Result: %.2f * %.2f = %.2f\n", c, d, mul(c, d))
	case 4:
		div(c, d)
	case 5:
		mod(c, d)
	default:
		fmt.Println("⚠️ Invalid choice! Please choose between 1–6.")
	}
	home()
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
		fmt.Println("⚠️ Error: division by zero‼️")
		return
	}
	fmt.Printf("Result: %.2f / %.2f = %.2f\n", c, d, c/d)
}

func mod(c, d float64) {
	if d == 0 {
		fmt.Println("⚠️ Error: modulus by zero‼️")
		return
	}
	fmt.Printf("Result: %.2f %% %.2f = %.2f\n", c, d, float64(int(c)%int(d)))
}
