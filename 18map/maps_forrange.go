package main

import "fmt"

func main() {
	map1 := make(map[int]int)
	map1[1] = 1
	map1[2] = 4
	map1[9] = 7
	map1[3] = 9

	for key, value := range map1 {
		fmt.Printf("key is %d,value is %d\n", key, value)
	}
}
