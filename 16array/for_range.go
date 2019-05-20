package main

import "fmt"

func main() {
	var arr1 [5]int

	for i, _ := range arr1 {
		arr1[i] = i * 4
	}
	for i, v := range arr1 {
		fmt.Printf("index %d is %d\n", i, v)
	}

}
