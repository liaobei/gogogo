package main

import "fmt"

func main() {
	ch := make(chan int)
	go tel(ch)
	for {
		v := <-ch
		if v == -1 {
			break
		}
		fmt.Printf("the value is %d", v)

	}

}

func tel(ch chan int) {
	for i := 0; i < 3; i++ {
		ch <- i
	}
	ch <- -1
}
