package main

import (
	"encoding/json"
	"fmt"
)

type data struct {
	Value int
	Left  *data
	Right *data
}

type bst struct {
	Root *data
}

func newBST() *bst {
	return &bst{}
}

func (b *bst) insert(value int) {
	data := &data{value, nil, nil}
	if b.isEmpty() {
		b.Root = data
		return
	}
	node, left, right, _ := b._getNode(value)
	if node.Value == value {
		fmt.Println("No dupes allowed")
		return
	}
	if left {
		node.Left = data
	}
	if right {
		node.Right = data
	}
}

func (b *bst) lookup(value int) (d *data) {
	if b.isEmpty() {
		return &data{}
	}
	node, _, _, _ := b._getNode(value)

	if node.Value != value {
		return &data{}
	}

	return node
}

func (b *bst) _getNode(value int) (d *data, left bool, right bool, previous *data) {
	finished := false

	currentNode := b.Root
	p := &data{}
	for !finished {
		right := currentNode.Value < value
		left := currentNode.Value > value

		if currentNode.Value == value {
			return currentNode, p.Left == currentNode, p.Right == currentNode, p
		}

		if right && currentNode.Right != nil {
			p = currentNode
			currentNode = currentNode.Right
		} else if right && currentNode.Right == nil {
			return currentNode, false, true, &data{}
		} else if left && currentNode.Left != nil {
			p = currentNode
			currentNode = currentNode.Left

		} else if left && currentNode.Left == nil {
			return currentNode, true, false, &data{}
		} else {
			finished = true
		}

	}

	return &data{}, false, false, &data{}

}

func (b *bst) remove(value int) {
	if b.isEmpty() {
		return
	}

	node, left, right, previous := b._getNode(value)

	if node.Left == nil && node.Right == nil {
		if left {
			previous.Left = nil
		}
		if right {
			previous.Right = nil
		}
	} else if node.Right == nil {
		n, p := b._traversalLeft(node, previous)
		fmt.Println(n)
		n.Left = node.Left
		n.Right = node.Right
		p.Left = nil
		if left {
			previous.Left = n
		}
		if right {
			previous.Right = n
		}
	} else if node.Left != nil {
		n, p := b._traversalLeft(node.Right, node)
		fmt.Println(n)
		p.Left = nil
		n.Left = node.Left
		n.Right = node.Right
		if left {
			previous.Left = n
		}
		if right {
			previous.Right = n
		}
	} else {

	}

	// previous 7
	// node.left 10
	// node.right8
	// if

}

func (b *bst) _traversalLeft(n *data, p *data) (node *data, previous *data) {
	if n.Left == nil {
		return n, p
	}

	return b._traversalLeft(n.Left, node)
}

func (b *bst) isEmpty() bool {
	if b.Root == nil {
		return true
	}
	return false
}

func main() {
	bst := newBST()
	fmt.Println(bst.isEmpty())
	bst.insert(9)
	bst.insert(4)
	bst.insert(6)
	bst.insert(20)
	bst.insert(170)
	bst.insert(170)
	bst.insert(15)
	bst.insert(1)
	// fmt.Println(bst.Root.Left.Left)
	fmt.Println()
	// fmt.Println(bst)
	bst.remove(20)
	b, _ := json.Marshal(bst)
	fmt.Println(string(b))

}
