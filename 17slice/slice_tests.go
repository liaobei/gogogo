package main

import (
	"bytes"
	"fmt"
)

func main() {
	//fmt.Println(mysplit("ABVDA", 2))
	//whatsres("qwertyuio")
	//fmt.Println(revert("Gor"))
	arr := [5]int{1, 2, 3, 4, 5}
	s := arr[:]
	fmt.Println(mapFunc(multiTen, s))
}

//练习7.12
func mysplit(str string, i int) (res1, res2 string) {
	s1 := []byte(str)[:i]
	s2 := []byte(str)[i:]
	res1 = string(s1)
	res2 = string(s2)
	return
}

//7.13
func whatsres(str string) {
	s := str[len(str)/2:] + str[:len(str)/2]
	fmt.Println(s)
}

//7.14
func revert(str string) (res string) {
	b := []byte(str)
	buf := bytes.Buffer{}
	for i := len(b) - 1; i >= 0; i-- {
		buf.WriteString(string(b[i]))
	}

	return buf.String()
}

//7.17
func mapFunc(fun func(i int) int, arr []int) []int {
	res := make([]int, 5)
	for i := 0; i < len(arr); i++ {
		res[i] = fun(arr[i])
	}
	return res
}
func multiTen(i int) int {
	res := i * 10
	return res
}

func function(a *int) {
	a = 123
}
