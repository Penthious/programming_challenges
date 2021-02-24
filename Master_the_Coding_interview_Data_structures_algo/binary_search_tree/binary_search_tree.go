package main

import "fmt"

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
			return currentNode, false, false, p
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

	// node, _, _, previous := b.lookup(value)

	// previous 7
	// node.left 10
	// node.right8
	// if

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
	fmt.Println(bst.Root.Left.Left)
	fmt.Println()
	fmt.Println(bst.lookup(30))

}
