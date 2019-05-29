package main

import "fmt"

type List []int

func (l List) len() int {
	return len(l)
}

func (l *List) append(val int) {
	*l = append(*l, val)
}

type Appender interface {
	append(int)
}

func CountInto(a Appender, start, end int) {
	for i := start; i <= end; i++ {
		a.append(i)
	}
}

type Lener interface {
	len() int
}

func LongEnough(l Lener) bool {
	return l.len()*10 > 42
}

func main() {
	var lst List

	if LongEnough(lst) {
		fmt.Printf("- lst is long enough\n")
	}

	plst := new(List)
	CountInto(plst, 1, 7)
	if LongEnough(plst) {
		fmt.Printf("- plst is long enough\n")
	}

}
