package main

func main() {
	var slice = make([]int, 10)

	for i := 0; i < cap(slice); i++ {
		slice[i] = i * i
	}
}
