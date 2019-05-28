package main

import "fmt"

type struct1 struct {
	i1  int
	f1  float32
	str string
}

type Inteval struct {
	start int
	end   int
}

type Point struct {
	x, y int
}

func main() {
	ms := new(struct1)
	ms.f1 = 10.5
	ms.i1 = 6
	ms.str = "liaobei"
	fmt.Printf("%v", ms)

	//intr := Inteval{0, 3}
	//intr := Inteval{end: 4, start: 1}
	//intr := Inteval{end: 5}
}
