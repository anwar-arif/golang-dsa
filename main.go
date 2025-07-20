package main

import (
	"fmt"
	"golang-dsa/priorityqueue"
)

func main() {
	pq := priorityqueue.NewMaxQueue(priorityqueue.StringCompare)
	pq.Push("apple")
	pq.Push("zebra")
	pq.Push("banana")

	for !pq.IsEmpty() {
		top, _ := pq.Pop()
		fmt.Println(top)
	}
}
