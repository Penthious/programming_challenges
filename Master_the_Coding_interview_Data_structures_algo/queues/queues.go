package main

import (
	"container/list"
)

func main() {
	queue := list.New()
	queue.PushBack(10)
	queue.PushBack(20)
	queue.PushBack(30)

	front := queue.Front()

	queue.Remove(front)

}
