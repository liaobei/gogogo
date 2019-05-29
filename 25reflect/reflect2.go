package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 3.444
	v := reflect.ValueOf(x)
	fmt.Println("settablility of v:", v.CanSet())
	v = reflect.ValueOf(&x)
	fmt.Println("type of v:", v.Type())
	fmt.Println("settability of v :", v.CanSet())
	v = v.Elem()
	fmt.Println("the elem of v is :", v)
	fmt.Println("settability of v:", v.CanSet())
	v.SetFloat(3.121)
	fmt.Println(v.Interface())
	fmt.Println(v)
}
