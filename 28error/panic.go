package main

import "fmt"

func main() {
	fmt.Println("start a program")
	panic("A severe error occurred:stopping the program")
	fmt.Println("Ending the program")
}
