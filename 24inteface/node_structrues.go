package main

import "fmt"

type Node struct {
	le   *Node
	data interface{}
	ri   *Node
}

func newNode(left, right *Node) *Node {
	return &Node{left, nil, right}
}
func (n *Node) SetData(data interface{}) {
	n.data = data
}
func main() {
	root := newNode(nil, nil)
	root.SetData("root node")

	a := newNode(nil, nil)
	a.SetData("left node")

	b := newNode(nil, nil)
	b.SetData("right node")

	root.le = a
	root.ri = b
	fmt.Printf("%v", root)
}
