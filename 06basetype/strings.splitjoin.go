package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "the quick brown fox jumps over the lazy dog"
	sl := strings.Fields(str)
	fmt.Printf("Splitted in slice:%v\n", sl)
	for _, val := range sl {
		fmt.Printf("%s -", val)
	}
	fmt.Println()
	fmt.Println("-----------")
	str2 := "GO1|THE ABC OF GO|25"
	sl2 := strings.Split(str2, "|")
	fmt.Printf("Splitted in slice:%V\n", sl2)
	for _, val := range sl2 {
		fmt.Printf("%s - ", val)
	}
	fmt.Println()
	str3 := strings.Join(sl2, ";")
	fmt.Println("sl2 joined by;:%s\n", str3)
}
