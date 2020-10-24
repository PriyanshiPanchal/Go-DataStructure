package main

import "fmt"

var count int

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

func (n *Node) Insert(k int) {
	if n.Key < k {
		if n.Right == nil {
			n.Right = &Node{Key: k}
		} else {
			n.Right.Insert(k)
		}
	} else if n.Key > k {
		if n.Left == nil {
			n.Left = &Node{Key: k}
		} else {
			n.Left.Insert(k)
		}
	}
}
func (n *Node) Search(k int) bool {

	count++

	if n == nil {
		return false
	}

	if n.Key < k {
		return n.Right.Search(k)
	} else if n.Key > k {
		return n.Left.Search(k)
	}

	return true
}
func main() {
	tree := &Node{Key: 100}
	tree.Insert(200)
	tree.Insert(300)
	tree.Insert(50)
	tree.Insert(53)
	tree.Insert(203)
	tree.Insert(34)
	tree.Insert(7)
	tree.Insert(250)
	tree.Insert(88)
	tree.Insert(24)
	tree.Insert(64)
	tree.Insert(276)
	fmt.Println(tree)

	fmt.Println(tree.Search(7))
	fmt.Println(count)
}
