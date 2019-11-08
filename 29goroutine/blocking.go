package main

//todo
import (
	"fmt"
	"time"
)

func f1(in chan int) {
	in <- 3
	in <- 31
	in <- 32
	fmt.Println(<-in)
}

func main() {
	out := make(chan int, 1)

	go f1(out)
	time.Sleep(6e10)

}
