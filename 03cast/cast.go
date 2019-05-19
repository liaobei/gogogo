package main

import "fmt"

func main() {
	type IZ int
	a := 5.0
	b := int(a)
	fmt.Print(b)

	//相同底层类型变量可以转换
	var c IZ = 5
	d := int(c)
	e := IZ(c)

}
