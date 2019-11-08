package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	go pump(ch1)
	go suck(ch1)
	time.Sleep(4 * time.Second)
	fmt.Println(<-ch1)
}
func suck(ch1 chan int) {
	for {
		fmt.Println(<-ch1)
	}
}

func pump(ch chan int) {
	for i := 0; ; i++ {
		ch <- i
	}
}
