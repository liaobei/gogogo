package main

import "fmt"

func main() {
	fv := func() {
		fmt.Println("Hello world")
	}
	fv()
	fmt.Printf("%T", fv)

}
