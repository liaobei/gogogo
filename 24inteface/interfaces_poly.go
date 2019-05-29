package main

import "fmt"

type Shaper1 interface {
	Area1() float32
}

type Rectangle struct {
	length, width float32
}

func (r *Rectangle) Area1() float32 {
	return r.length * r.width
}

func main() {
	r := &Rectangle{5, 4}
	shapes := []Shaper1{r}
	for n, _ := range shapes {
		fmt.Println(shapes[n].Area1())
	}

}
