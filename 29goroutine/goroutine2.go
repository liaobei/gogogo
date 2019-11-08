package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan string)
	go sendData(ch)
	go getData(ch)
	time.Sleep(10e9)
	defer close(ch)

}
func getData(ch chan string) {
	var input string
	for {
		input = <-ch
		fmt.Printf("%s\n", input)
	}
}
func sendData(ch chan string) {
	ch <- "jiangxi"
	time.Sleep(5 * time.Second)
	ch <- "beijing"
	ch <- "hangzhou"
	ch <- "shenzheng"
}
