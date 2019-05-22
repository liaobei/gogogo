package main

import (
	"fmt"
	"sort"
)

var (
	varVal = map[string]int{"alpha": 34, "bravo": 56, "charlie": 23,
		"delta": 87, "echo": 56, "foxtrot": 12,
		"golf": 34, "hotel": 16, "indio": 87,
		"juliet": 65, "kili": 43, "lima": 98}
)

func main() {
	fmt.Println("not sort")
	for k, v := range varVal {
		fmt.Printf("key is %s,value is %d\n", k, v)
	}
	keys := make([]string, len(varVal))
	i := 0
	for k, _ := range varVal {
		keys[i] = k
		i++
	}

	sort.Strings(keys)
	fmt.Println()
	fmt.Println("sort")
	for _, k := range keys {
		fmt.Printf("Key: %v, Value: %v\n ", k, varVal[k])
	}
}
