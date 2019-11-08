package main

import (
	"fmt"
	"time"
)

func main() {
	steam := pump1()
	go suck1(steam)
	time.Sleep(1e9)
}

func pump1() chan int {
	ch := make(chan int)
	go func() {
		for i := 0; ; i++ {
			ch <- i
		}
	}()
	return ch
}

func suck1(ch chan int) {
	for {
		fmt.Println(<-ch)
	}
}
