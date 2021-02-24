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
	node, left, right := b.lookup(value)
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

func (b *bst) lookup(value int) (d *data, left bool, right bool) {
	if b.isEmpty() {
		return &data{}, false, false
	}
	finished := false

	currentNode := b.Root
	for !finished {
		right := currentNode.Value < value
		left := currentNode.Value > value

		if currentNode.Value == value {
			return currentNode, false, false
		}

		if right && currentNode.Right != nil {
			fmt.Println(b.Root)
			currentNode = currentNode.Right
		} else if right && currentNode.Right == nil {
			return currentNode, false, true
		} else if left && currentNode.Left != nil {
			currentNode = currentNode.Left

		} else if left && currentNode.Left == nil {
			return currentNode, true, false
		} else {
			finished = true
		}

	}

	return &data{}, false, false
}

func (b *bst) remove(value int) {}

func (b *bst) isEmpty() bool {
	if b.Root == nil {
		return true
	}
	return false
}

func main() {
	bst := newBST()
	fmt.Println(bst.isEmpty())
	bst.insert(10)
	fmt.Println(bst.isEmpty())
	bst.insert(15)
	bst.insert(13)
	bst.insert(14)
	bst.insert(14)
	fmt.Println(bst.Root.Right.Left.Right)
	fmt.Println(bst.lookup(14))

}
