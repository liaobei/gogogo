package main

import (
	"fmt"
	"reflect"
)

type NotknownType struct {
	S1, S2, S3 string
}

func (n NotknownType) String() string {
	return n.S1 + " - " + n.S2 + " - " + n.S3
}

// variable to investigate:
//var secret interface{} =

func main() {
	secret := NotknownType{"Ada", "Go", "Oberon"}
	value := reflect.ValueOf(&secret).Elem() // <main.NotknownType Value>
	typ := value.Type()                      // main.NotknownType
	// alternative:
	//typ := value.Type()  // main.NotknownType
	fmt.Println(typ)
	knd := value.Kind() // struct
	fmt.Println(knd)

	// iterate through the fields of the struct:
	for i := 0; i < value.NumField(); i++ {
		//f:=value.Field(i)
		fmt.Printf("Field %d: %s\n", i, typ.Field(i).Name)
		// error: panic: reflect.Value.SetString using value obtained using unexported field
		value.Field(i).SetString("C#")
	}

	// call the first method, which is String():
	results := value.Method(0).Call(nil)
	fmt.Println(results) // [Ada - Go - Oberon]
}
