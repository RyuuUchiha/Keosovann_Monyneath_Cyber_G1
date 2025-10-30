package main

import "fmt"

func main() {

	header()
	lab0()
	header()
	lab1()
	header()
	lab2()
	header()
	lab3()
	header()
}

func header() {
	for i := 0; i < 156; i++ {
		fmt.Print("-")
	}
	fmt.Println()
}
