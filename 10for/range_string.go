package main

import "fmt"

func main() {
	str := "go is beautiful language"
	for pos, char := range str {
		fmt.Printf("character on position %d is :%c \n", pos, char)
	}
}
