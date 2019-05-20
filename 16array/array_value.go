package main

import "fmt"

//证明当数组赋值时，发生了数组内存拷贝
func main() {
	var arr = new([3]int)
	arr[0] = 1
	//值拷贝
	//arr1 := *arr
	//内存拷贝
	arr1 := arr
	arr1[0] = 2
	fmt.Print(arr[0])

}
