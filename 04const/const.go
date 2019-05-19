package main

func main() {
	//常量必须在编译器就能够确定
	const PI = 3.14159
	// 错误写法
	//const c1 = getNumber()   getNumber userd as value

	const (
		unKnow = 0
		Female = 1
		Mal2   = 2
	)

	//第一个iota等于0，每当新一行被用将会递增
	const (
		a = iota
		b
		c
	)
}
