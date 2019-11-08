package main

var send_only chan<- int
var recv_only <-chan int

func main() {
	var c = make(chan int)
	go source2(c)
	go sink2(c)
}

func source2(ch chan<- int) {
	for {
		ch <- 1
	}
}
func sink2(ch <-chan int) {
	for {
		<-ch
	}
}
