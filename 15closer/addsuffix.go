package main

import (
	"fmt"
	"strings"
)

func main() {
	addBmp := addsuffix("bmp")
	fmt.Println(addBmp(("hh")))
}

func addsuffix(suffix string) func(string) string {

	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + "." + suffix
		}
		return name
	}
}
