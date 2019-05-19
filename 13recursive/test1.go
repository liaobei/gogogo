package main

import "fmt"

func main() {
	num := 1
	myPrint(num)
}
func myPrint(i int) {
	if i == 10 {
		fmt.Println(i)
		return
	}
	myPrint(i + 1)
	fmt.Println(i)
}
