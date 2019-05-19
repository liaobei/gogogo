package main

import "fmt"

func main() {
	sum := 0
	func() {

		for i := 1; i < 1e6; i++ {
			sum += 1
		}
	}()
	fmt.Println(sum)
}
