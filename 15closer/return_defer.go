package main

func ff() (ret int) {
	defer func() {
		ret++
	}()
	return 1
}
func main() {
	println(ff())
}
