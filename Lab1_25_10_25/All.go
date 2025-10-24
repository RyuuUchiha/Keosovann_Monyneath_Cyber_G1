package main

import "fmt"


func main() {

	Header()
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
	Header()
	
}


func Header() {
	for i := 0; i < 50; i++ {
		fmt.Print("-")
	}
	fmt.Println()
}

