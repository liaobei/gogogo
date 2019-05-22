package main

import "fmt"

func main() {
	var value int
	var isPresent bool

	map1 := make(map[string]int)
	map1["new Delhi"] = 55
	map1["beiJing"] = 20
	map1["jiangxi"] = 30

	if value, isPresent = map1["beiJing"]; isPresent {
		delete(map1, "beijing")

		fmt.Printf("%d,OK", value)
	} else {
		fmt.Println("error")
	}

}
