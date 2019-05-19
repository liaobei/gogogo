package main

import "fmt"

func main() {
	var i1 int
	var f1 float32
	_, _, f1 = ThreeValues()
	fmt.Println(i1, f1)
}
func ThreeValues() (int, int, float32) {
	return 5, 6, 7.5
}
