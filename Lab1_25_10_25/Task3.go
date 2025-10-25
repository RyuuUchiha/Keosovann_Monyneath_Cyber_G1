package main

import (
	"bufio"
	"fmt"
	"os"
)

func task3() {

	fmt.Println("	📋Task 3: Bitwise and Assignment Opr.")

	var a, b int

	fmt.Print("Enter first integer(a): ")
	if _, err := fmt.Scan(&a); err != nil {
		fmt.Println("❌ Invalid input! Please enter a number only.")
		bufio.NewReader(os.Stdin).ReadString('\n')
		return
	}

	fmt.Print("Enter second integer(b): ")
	if _, err := fmt.Scan(&b); err != nil {
		fmt.Println("❌ Invalid input! Please enter a number only.")
		bufio.NewReader(os.Stdin).ReadString('\n')
		return
	}

	fmt.Println("\n--- Bitwise Operations ---")
	myAND(a, b)
	myOr(a, b)
	myXOR(a, b)
	myNot(a, b)
	myLeftShift(a, 1)
	myRightShift(a, 1)

	fmt.Println("\n--- Assignment Operators ---")
	num := a
	fmt.Printf("Initial value: num = a = %d\n", num)

	num &= b
	fmt.Printf("num &= %d → %d\n", b, num)

	num |= b
	fmt.Printf("num |= %d → %d\n", b, num)

	num ^= b
	fmt.Printf("num ^= %d → %d\n", b, num)

	num <<= 1
	fmt.Printf("num <<= 1 → %d\n", num)

	num >>= 1
	fmt.Printf("num >>= 1 → %d\n", num)
}

func myAND(a, b int) {
	fmt.Printf("Bitwise AND: %d & %d = %d\n", a, b, a&b)
}

func myOr(a, b int) {
	fmt.Printf("Bitwise OR: %d | %d = %d\n", a, b, a|b)
}

func myXOR(a, b int) {
	fmt.Printf("Bitwise XOR: %d ^ %d = %d\n", a, b, a^b)
}

func myNot(a, b int) {
	fmt.Printf("Bitwise NOT: ^%d = %d\n", a, ^a)
	fmt.Printf("Bitwise NOT: ^%d = %d\n", b, ^b)
}

func myLeftShift(a int, n int) {
	fmt.Printf("%d << %d = %d\n", a, n, a<<n)
}

func myRightShift(a int, n int) {
	fmt.Printf("%d >> %d = %d\n", a, n, a>>n)
}
