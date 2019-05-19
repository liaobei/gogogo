package main

import "fmt"

func MinMax(a int, b int) (min int, max int) {
	if a < b {
		min = a
		max = b
	} else {
		min = b
		max = a
	}
	return
}

func main() {
	var min, max int

	min, max = MinMax(55, 66)
	fmt.Print(min, max)
}
