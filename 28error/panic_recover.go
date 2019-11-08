package main

import "fmt"

func badCall() {
	panic("bad end")
}
func test() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("panicing %s \r\n", e)
		}
	}()
	badCall()
	fmt.Println("after bad Call")
}

func main() {
	fmt.Printf("Calling test\r\n")
	test()
	fmt.Println("Test completed\r\n")
}
