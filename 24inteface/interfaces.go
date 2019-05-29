package main

import "fmt"

type Shaper interface {
	Area() float32
}
type Squaer struct {
	side float32
}

func (sq *Squaer) Area() float32 {
	return sq.side * sq.side
}

func main() {
	sq1 := new(Squaer)
	sq1.side = 5
	var areaIntf Shaper
	areaIntf = sq1

	fmt.Printf("The square has area: %f\n", areaIntf.Area())
}
