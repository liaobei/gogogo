package main

func main() {
	var arr = [5]int{1, 2, 3, 4, 5}
	sum(arr[:])
}
func sum(a []int) (s int) {
	s = 0
	for i := 0; i < len(a); i++ {
		s += a[i]
	}
	return
}
