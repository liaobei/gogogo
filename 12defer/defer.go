package main

import "fmt"

func main() {
	function1()
}
func function1() {
	fmt.Println("In function1 at the top\n")
	defer function2()
	fmt.Println("In function1 at the bottom")
}
func function2() {
	fmt.Println("Function2 defer until the end of the calling function!")
}
