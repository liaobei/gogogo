package main

import "fmt"

var (
	firstName, lastName, s string
	i                      int
	f                      float32
	input                  = "56.13 /4334 /GO"
	fomate                 = "%f / %d / %s"
)

func main() {
	fmt.Println("Please enter your full name: ")
	fmt.Scanln(&firstName, &lastName)

	fmt.Printf("Hi %s %s!\n", firstName, lastName)

}
