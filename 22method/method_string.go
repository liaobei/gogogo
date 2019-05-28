package main

import (
	"fmt"
	"strconv"
)

type TwoIntsz struct {
	a int
	b int
}

func main() {
	two1 := new(TwoIntsz)
	two1.a = 12
	two1.b = 10
	fmt.Printf("two1 is: %v\n", two1)
	fmt.Println("two1 is:", two1)
	//完全规格
	fmt.Printf("two1 is: %T\n", two1)
	fmt.Printf("two1 is: %#v\n", two1)
}

func (tn *TwoIntsz) String() string {
	return "(" + strconv.Itoa(tn.a) + "/" + strconv.Itoa(tn.b) + ")"

}
