package main

import "fmt"

type SimpleI interface {
	get() int
	set()
}

type Simple struct {
}

func (s *Simple) get() (i int) {
	i = 1
	return
}
func (s *Simple) set() {
	fmt.Printf("set")
}

func main() {
	s := new(Simple)
	var ss SimpleI
	ss = s
	ss.get()
	ss.set()
}
