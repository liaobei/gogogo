package main

import (
	"fmt"
	"reflect"
)

type tagType struct {
	field1 bool   "An important answer"
	field2 string "The name of the thing"
	field3 int    "How much there are"
}

func refTag(tt tagType, ix int) {
	ttType := reflect.TypeOf(tt)
	idxField := ttType.Field(ix)
	fmt.Printf("%v\n", idxField.Tag)
}
func main() {
	tt := tagType{true, "liaobei", 399}
	for i := 0; i < 3; i++ {
		refTag(tt, i)
	}
}
