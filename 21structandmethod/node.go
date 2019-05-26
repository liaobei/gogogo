package main

type Node struct {
	pr   *Node
	data float64
	su   *Node
}

type Tree struct {
	le   *Tree
	data float64
	ri   *Tree
}
