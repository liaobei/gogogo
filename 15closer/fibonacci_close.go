package main

import "fmt"

func main() {
	f := fibo()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

func fibo() func() int {
	a := 0
	b := 1
	return func() int {
		c := a
		a, b = b, b+c
		return c
	}
}
