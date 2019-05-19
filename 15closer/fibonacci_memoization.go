package main

import (
	"fmt"
	"time"
)

const LIM = 41

var fibs [LIM]uint64

func main() {
	var result uint64 = 0
	start := time.Now()
	for i := 0; i < LIM; i++ {
		result = fibonacci(i)
		fmt.Printf("fibonacci(%d) is:%d\n", i, result)
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Println(delta)

}
func fibonacci(i int) (res uint64) {
	if fibs[i] != 0 {
		res = fibs[i]
		return
	}
	if i <= 1 {
		res = 1
	} else {
		res = fibonacci(i-1) + fibonacci(i-2)
	}
	fibs[i] = res
	return
}
