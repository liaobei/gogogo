package main

import (
	"errors"
	"fmt"
)

var errorNotFound error = errors.New("Not found error")

func main() {
	fmt.Printf("error :%v", errorNotFound)
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("Math - square root of nagative numver")
	}
	return 0.1, errors.New("Math - square root of nagative numver")
}
