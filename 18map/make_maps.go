package main

func main() {
	var mapLit map[string]int
	var mapAssigned map[string]int
	mapLit = map[string]int{"one": 1, "two": 2}
	mapCreated := make(map[string]float32)
	mapAssigned = mapLit
	mapCreated["key"] = 4.5
	mapCreated["key2"] = 4.6
	mapAssigned["two"] = 3
}
