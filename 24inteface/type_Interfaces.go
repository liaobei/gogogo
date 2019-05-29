package main

import (
	"fmt"
	"math"
)

type Sqnare struct {
	side float32
}

type Cirrcle struct {
	radius float32
}
type Shabber interface {
	Area2() float32
}

func main() {
	var areaInf Shabber
	sq1 := new(Sqnare)
	sq1.side = 5
	areaInf = sq1
	if _, ok := areaInf.(*Sqnare); ok {
		fmt.Printf("okk")
	}
	if _, ok := areaInf.(*Cirrcle); ok {
		fmt.Printf("okkk")
	} else {
		fmt.Printf("no okk")
	}

}

func (sq *Sqnare) Area2() float32 {
	return sq.side * sq.side
}

func (ci *Cirrcle) Area2() float32 {
	return ci.radius * ci.radius * math.Pi
}
