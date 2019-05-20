package main

import (
	"bytes"
	"fmt"
)

func main() {
	var buffer bytes.Buffer
	buffer.WriteString("aaaa")
	buffer.WriteString("bbb")
	aa := buffer.Bytes()[3:]
	fmt.Print(aa)
}
