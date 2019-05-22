package main

import "fmt"

func main() {
	mf := map[int]func() int{
		1: func() int { return 10 },
		2: func() int { return 20 },
		3: func() int { return 30 },
	}

	//map2:=make(map[string]float32,100)
	fmt.Println(mf[1]())
}
