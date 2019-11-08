package main

import "fmt"

func dailyTemperatures(T []int) []int {
	var length = len(T)
	var res = make([]int, length)
	for i := 0; i < length-1; i++ {
		for j := i + 1; j < length; j++ {
			if T[j] > T[i] {
				res[i] = j - i
				break
			}
		}
	}
	return res
}
func main() {
	var T = []int{73, 74, 75, 71, 69, 72, 76, 73}
	fmt.Println(dailyTemperatures(T))
}
