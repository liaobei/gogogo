package main

import (
	"fmt"
	"strings"
)

func main() {
	var str1 string = "this is an example of a string"

	fmt.Printf("T/F Does the string \"%s\"have a prefix %s\n", str1, "TH")
	fmt.Printf("%t\n", strings.HasPrefix(str1, "TH"))

	var str string = "Hi, I'm Marc, Hi."

	fmt.Printf("The position of \"Marc\" is: ")
	fmt.Printf("%d\n", strings.Index(str, "Marc"))

	fmt.Printf("The position of the first instance of \"Hi\" is: ")
	fmt.Printf("%d\n", strings.Index(str, "Hi"))
	fmt.Printf("The position of the last instance of \"Hi\" is: ")
	fmt.Printf("%d\n", strings.LastIndex(str, "Hi"))

	fmt.Printf("The position of \"Burger\" is: ")
	fmt.Printf("%d\n", strings.Index(str, "Burger"))
}
