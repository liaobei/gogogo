package main

import (
	"fmt"
	"strconv"
)

func main() {
	org := "ABC"
	var newS string

	fmt.Printf("The size int is :%d\n", strconv.IntSize)

	an, errr := strconv.Atoi(org)

	if errr != nil {
		fmt.Printf("%s", errr)
		fmt.Printf("org %s isnot a integer -exiting with error\n", org)
		return
	}

	fmt.Printf("The Integer is %d\n", an)
	an += 5
	newS = strconv.Itoa(an)
	fmt.Printf("The new string is :%s\n", newS)

}
