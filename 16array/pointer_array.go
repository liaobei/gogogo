package main

import "fmt"

func main() {
	var arr1 = new([5]int) // *[5]int
	var arr2 [5]int        //  [5]int
	arr2 = *arr1
	arr2[0] = 1
	//fmt.Print(arr1[0]) //0

	var ar [3]int
	f(ar)
	fp(&ar)

}
func fp(ints *[3]int) {
	fmt.Println(ints)
}
func f(ints [3]int) {
	fmt.Println(ints)
}
