package main

import "fmt"

func fpp(a *[3]int) { fmt.Print(a) }

func main() {

	for i := 0; i < 3; i++ {
		fpp(&[3]int{i, i * i, i * i * i})
	}
}
