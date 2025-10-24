package main

import "fmt"


func main() {

	task1()
	Header()
	task2()
	Header()
	task3()
	Header()
	task4()
	Header()
	task5()
	Header()
	task6()
	
}


func Header() {
	for i := 0; i < 50; i++ {
		fmt.Print("-")
	}
	fmt.Println()
}

