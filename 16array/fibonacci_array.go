package main

import "fmt"

func main() {
	var fibs [50]int
	fibs[0] = 0
	fibs[1] = 1
	for i := 2; i < 50; i++ {
		fibs[i] = fibs[i-1] + fibs[i-2]
	}
	fmt.Println(fibs)
}
