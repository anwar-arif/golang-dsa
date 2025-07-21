// Package queue provides a generic FIFO (First In, First Out) queue implementation.
package queue

import (
	"fmt"
)

// Node represents a node in the queue
type Node[T any] struct {
	Value T
	Next  *Node[T]
}

// Queue represents a FIFO queue
type Queue[T any] struct {
	front *Node[T] // Points to the first element (dequeue from here)
	rear  *Node[T] // Points to the last element (enqueue to here)
	size  int
}

// NewQueue creates a new empty queue
func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		front: nil,
		rear:  nil,
		size:  0,
	}
}

// Push adds an item to the rear of the queue
func (q *Queue[T]) Push(value T) {
	newNode := &Node[T]{
		Value: value,
		Next:  nil,
	}

	if q.IsEmpty() {
		// First element
		q.front = newNode
		q.rear = newNode
	} else {
		// Add to rear
		q.rear.Next = newNode
		q.rear = newNode
	}

	q.size++
}

// Pop removes and returns the item from the front of the queue
func (q *Queue[T]) Pop() (T, error) {
	var zero T

	if q.IsEmpty() {
		return zero, fmt.Errorf("queue is empty")
	}

	value := q.front.Value
	q.front = q.front.Next

	// If queue becomes empty, reset rear as well
	if q.front == nil {
		q.rear = nil
	}

	q.size--
	return value, nil
}

// Front returns the front item without removing it
func (q *Queue[T]) Front() (T, error) {
	var zero T

	if q.IsEmpty() {
		return zero, fmt.Errorf("queue is empty")
	}

	return q.front.Value, nil
}

// Rear returns the rear item without removing it
func (q *Queue[T]) Rear() (T, error) {
	var zero T

	if q.IsEmpty() {
		return zero, fmt.Errorf("queue is empty")
	}

	return q.rear.Value, nil
}

// IsEmpty returns true if the queue is empty
func (q *Queue[T]) IsEmpty() bool {
	return q.front == nil
}

// Size returns the number of items in the queue
func (q *Queue[T]) Size() int {
	return q.size
}

// Clear removes all items from the queue
func (q *Queue[T]) Clear() {
	q.front = nil
	q.rear = nil
	q.size = 0
}

// ToSlice returns all items as a slice from front to rear
func (q *Queue[T]) ToSlice() []T {
	result := make([]T, 0, q.size)
	current := q.front

	for current != nil {
		result = append(result, current.Value)
		current = current.Next
	}

	return result
}

// String returns a string representation of the queue
func (q *Queue[T]) String() string {
	return fmt.Sprintf("Queue{size: %d, front->rear: %v}", q.size, q.ToSlice())
}

// Example usage and demonstrations
func ExampleUsage() {
	fmt.Println("=== Generic Queue Examples ===\n")

	// Example 1: Basic integer queue
	fmt.Println("1. Basic Integer Queue (FIFO):")
	intQueue := NewQueue[int]()

	intQueue.Push(10)
	intQueue.Push(20)
	intQueue.Push(30)

	fmt.Printf("Queue size: %d\n", intQueue.Size())
	fmt.Printf("Front item: ")
	if front, err := intQueue.Front(); err == nil {
		fmt.Printf("%d\n", front)
	}

	fmt.Println("Dequeuing items:")
	for !intQueue.IsEmpty() {
		val, _ := intQueue.Pop()
		fmt.Printf("  Dequeued: %d\n", val)
	}

	// Example 2: String queue
	fmt.Println("\n2. String Queue:")
	stringQueue := NewQueue[string]()

	stringQueue.Push("first")
	stringQueue.Push("second")
	stringQueue.Push("third")

	fmt.Println("Queue contents:", stringQueue.ToSlice())

	for !stringQueue.IsEmpty() {
		val, _ := stringQueue.Pop()
		fmt.Printf("  Dequeued: %s\n", val)
	}

	// Example 3: Custom type queue
	fmt.Println("\n3. Custom Type Queue (Tasks):")

	type Task struct {
		ID   int
		Name string
	}

	taskQueue := NewQueue[Task]()

	taskQueue.Push(Task{ID: 1, Name: "First task"})
	taskQueue.Push(Task{ID: 2, Name: "Second task"})
	taskQueue.Push(Task{ID: 3, Name: "Third task"})

	fmt.Println("Processing tasks in FIFO order:")
	for !taskQueue.IsEmpty() {
		task, _ := taskQueue.Pop()
		fmt.Printf("  Processing: Task %d - %s\n", task.ID, task.Name)
	}

	// Example 4: Queue operations
	fmt.Println("\n4. Queue Operations:")
	opQueue := NewQueue[string]()

	opQueue.Push("A")
	opQueue.Push("B")
	opQueue.Push("C")

	fmt.Println("Original queue:", opQueue.ToSlice())

	front, _ := opQueue.Front()
	rear, _ := opQueue.Rear()
	fmt.Printf("Front: %s, Rear: %s\n", front, rear)
	fmt.Println()

	// Example 6: BFS-like usage
	fmt.Println("\n6. BFS-like Usage (Processing Levels):")

	type Node struct {
		Value    string
		Children []string
	}

	bfsQueue := NewQueue[Node]()

	// Simulate a tree traversal
	root := Node{Value: "Root", Children: []string{"A", "B", "C"}}
	bfsQueue.Push(root)

	level := 0
	fmt.Printf("Level %d: %s\n", level, root.Value)
	level++

	// Process children (simplified BFS)
	for _, child := range root.Children {
		childNode := Node{Value: child, Children: []string{}}
		bfsQueue.Push(childNode)
	}

	fmt.Printf("Level %d: ", level)
	for !bfsQueue.IsEmpty() {
		node, _ := bfsQueue.Pop()
		fmt.Printf("%s ", node.Value)
	}
	fmt.Println()

	// Example 7: Error handling
	fmt.Println("\n7. Error Handling:")
	emptyQueue := NewQueue[int]()

	_, err := emptyQueue.Pop()
	if err != nil {
		fmt.Printf("Pop error: %v\n", err)
	}

	_, err = emptyQueue.Front()
	if err != nil {
		fmt.Printf("Front error: %v\n", err)
	}

	_, err = emptyQueue.Rear()
	if err != nil {
		fmt.Printf("Rear error: %v\n", err)
	}
}
