package main

import "fmt"

func main() {
	slice1 := make([]int, 0, 10)
	for i := 0; i < cap(slice1); i++ {
		slice1 = slice1[0 : i+1]
		slice1[i] = i
		fmt.Printf("The length of slice is %d\n", len(slice1))
	}

	var ar = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	a1 := ar[5:7]
	fmt.Println(a1)
	a1 = a1[1:5]
	fmt.Println(a1)

	println(len(a1[1:2]))

}
