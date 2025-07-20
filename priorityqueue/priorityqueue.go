package priorityqueue

import (
	"container/heap"
	"fmt"
)

// CompareFunc defines a comparison function type
// Returns:
//
//	-1 if a < b
//	 0 if a == b
//	 1 if a > b
type CompareFunc[T any] func(a, b T) int

// Item represents an item in the priority queue
type Item[T any] struct {
	Value T
	Index int // internal index for heap operations
}

// NewItem creates a new item with value
func NewItem[T any](value T) *Item[T] {
	return &Item[T]{
		Value: value,
	}
}

// internal heap implementation
type priorityHeap[T any] struct {
	items     []*Item[T]
	compare   CompareFunc[T]
	isMaxHeap bool
}

func (h *priorityHeap[T]) Len() int { return len(h.items) }

func (h *priorityHeap[T]) Less(i, j int) bool {
	cmp := h.compare(h.items[i].Value, h.items[j].Value)
	if h.isMaxHeap {
		return cmp > 0 // For max-heap, reverse the comparison
	}
	return cmp < 0 // For min-heap, use normal comparison
}

func (h *priorityHeap[T]) Swap(i, j int) {
	h.items[i], h.items[j] = h.items[j], h.items[i]
	h.items[i].Index = i
	h.items[j].Index = j
}

func (h *priorityHeap[T]) Push(x interface{}) {
	item := x.(*Item[T])
	item.Index = len(h.items)
	h.items = append(h.items, item)
}

func (h *priorityHeap[T]) Pop() interface{} {
	old := h.items
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.Index = -1 // for safety
	h.items = old[0 : n-1]
	return item
}

// PriorityQueue represents a priority queue with custom comparison
type PriorityQueue[T any] struct {
	heap *priorityHeap[T]
}

// NewMinQueue creates a new min-priority queue using the provided compare function
// Items that compare as "less" will have higher priority
func NewMinQueue[T any](compare CompareFunc[T]) *PriorityQueue[T] {
	h := &priorityHeap[T]{
		items:     make([]*Item[T], 0),
		compare:   compare,
		isMaxHeap: false,
	}
	heap.Init(h)
	return &PriorityQueue[T]{heap: h}
}

// NewMaxQueue creates a new max-priority queue using the provided compare function
// Items that compare as "greater" will have higher priority
func NewMaxQueue[T any](compare CompareFunc[T]) *PriorityQueue[T] {
	h := &priorityHeap[T]{
		items:     make([]*Item[T], 0),
		compare:   compare,
		isMaxHeap: true,
	}
	heap.Init(h)
	return &PriorityQueue[T]{heap: h}
}

// Push adds an item to the priority queue
func (pq *PriorityQueue[T]) Push(value T) {
	item := NewItem(value)
	heap.Push(pq.heap, item)
}

// Pop removes and returns the item with highest priority
func (pq *PriorityQueue[T]) Pop() (T, error) {
	var zero T
	if pq.IsEmpty() {
		return zero, fmt.Errorf("priority queue is empty")
	}
	item := heap.Pop(pq.heap).(*Item[T])
	return item.Value, nil
}

// Peek returns the item with highest priority without removing it
func (pq *PriorityQueue[T]) Peek() (T, error) {
	var zero T
	if pq.IsEmpty() {
		return zero, fmt.Errorf("priority queue is empty")
	}
	return pq.heap.items[0].Value, nil
}

// IsEmpty returns true if the priority queue is empty
func (pq *PriorityQueue[T]) IsEmpty() bool {
	return pq.heap.Len() == 0
}

// Size returns the number of items in the priority queue
func (pq *PriorityQueue[T]) Size() int {
	return pq.heap.Len()
}

// UpdateItem triggers a re-heapify for an item after it has been modified
// You should modify the item externally, then call this method
func (pq *PriorityQueue[T]) UpdateItem(item *Item[T]) {
	heap.Fix(pq.heap, item.Index)
}

// Remove removes an item from the priority queue
func (pq *PriorityQueue[T]) Remove(item *Item[T]) {
	heap.Remove(pq.heap, item.Index)
}

// ToSlice returns all items as a slice (does not modify the queue)
func (pq *PriorityQueue[T]) ToSlice() []*Item[T] {
	result := make([]*Item[T], len(pq.heap.items))
	copy(result, pq.heap.items)
	return result
}

// Clear removes all items from the priority queue
func (pq *PriorityQueue[T]) Clear() {
	pq.heap.items = pq.heap.items[:0]
	heap.Init(pq.heap)
}

// String returns a string representation of the priority queue
func (pq *PriorityQueue[T]) String() string {
	return fmt.Sprintf("PriorityQueue{size: %d}", pq.Size())
}

// Common comparison functions

// IntCompare compares two integers
func IntCompare(a, b int) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return 0
}

// StringCompare compares two strings lexicographically
func StringCompare(a, b string) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return 0
}

// Float64Compare compares two float64 values
func Float64Compare(a, b float64) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return 0
}

// ReverseCompare reverses any comparison function
func ReverseCompare[T any](compare CompareFunc[T]) CompareFunc[T] {
	return func(a, b T) int {
		return -compare(a, b)
	}
}

// Example types and their compare functions

// Task represents a task with priority
type Task struct {
	ID       int
	Name     string
	Priority int
}

// TaskByPriority compares tasks by their priority (lower number = higher priority)
func TaskByPriority(a, b Task) int {
	return IntCompare(a.Priority, b.Priority)
}

// TaskByID compares tasks by their ID
func TaskByID(a, b Task) int {
	return IntCompare(a.ID, b.ID)
}

// Node represents a graph node with distance
type Node struct {
	ID       int
	Distance int
}

// NodeByDistance compares nodes by their distance
func NodeByDistance(a, b Node) int {
	return IntCompare(a.Distance, b.Distance)
}

// Patient represents a hospital patient
type Patient struct {
	Name         string
	Age          int
	UrgencyLevel int
}

// PatientByUrgency compares patients by urgency level (higher = more urgent)
func PatientByUrgency(a, b Patient) int {
	return IntCompare(a.UrgencyLevel, b.UrgencyLevel)
}

// Score represents a game score
type Score struct {
	PlayerName string
	Points     int
}

// ScoreByPoints compares scores by points
func ScoreByPoints(a, b Score) int {
	return IntCompare(a.Points, b.Points)
}

// Example usage demonstrating different scenarios
func ExampleUsage() {
	fmt.Println("=== Generic Priority Queue Examples ===")

	// Example 1: Integer min-heap
	fmt.Println("1. Integer Min-Heap:")
	intMinQueue := NewMinQueue(IntCompare)

	intMinQueue.Push(30)
	intMinQueue.Push(10)
	intMinQueue.Push(20)

	for !intMinQueue.IsEmpty() {
		val, _ := intMinQueue.Pop()
		fmt.Printf("  Popped: %d\n", val)
	}

	// Example 2: String max-heap
	fmt.Println("\n2. String Max-Heap:")
	stringMaxQueue := NewMaxQueue(StringCompare)

	stringMaxQueue.Push("apple")
	stringMaxQueue.Push("zebra")
	stringMaxQueue.Push("banana")

	for !stringMaxQueue.IsEmpty() {
		val, _ := stringMaxQueue.Pop()
		fmt.Printf("  Popped: %s\n", val)
	}

	// Example 3: Task scheduling (min-heap by priority)
	fmt.Println("\n3. Task Scheduling (Min-Heap by Priority):")
	taskQueue := NewMinQueue(TaskByPriority)

	taskQueue.Push(Task{ID: 1, Name: "Send email", Priority: 3})
	taskQueue.Push(Task{ID: 2, Name: "Fix critical bug", Priority: 1})
	taskQueue.Push(Task{ID: 3, Name: "Update docs", Priority: 2})

	for !taskQueue.IsEmpty() {
		task, _ := taskQueue.Pop()
		fmt.Printf("  Processing: %s (Priority: %d)\n", task.Name, task.Priority)
	}

	// Example 4: Dijkstra's algorithm (min-heap by distance)
	fmt.Println("\n4. Dijkstra's Algorithm (Min-Heap by Distance):")
	dijkstraQueue := NewMinQueue(NodeByDistance)

	dijkstraQueue.Push(Node{ID: 0, Distance: 0})
	dijkstraQueue.Push(Node{ID: 1, Distance: 4})
	dijkstraQueue.Push(Node{ID: 2, Distance: 2})
	dijkstraQueue.Push(Node{ID: 3, Distance: 6})

	for !dijkstraQueue.IsEmpty() {
		node, _ := dijkstraQueue.Pop()
		fmt.Printf("  Visit node %d (Distance: %d)\n", node.ID, node.Distance)
	}

	// Example 5: Hospital emergency room (max-heap by urgency)
	fmt.Println("\n5. Hospital Emergency Room (Max-Heap by Urgency):")
	hospitalQueue := NewMaxQueue(PatientByUrgency)

	hospitalQueue.Push(Patient{Name: "John", Age: 25, UrgencyLevel: 3})
	hospitalQueue.Push(Patient{Name: "Sarah", Age: 45, UrgencyLevel: 8})
	hospitalQueue.Push(Patient{Name: "Mike", Age: 30, UrgencyLevel: 1})

	fmt.Println("  Treatment order:")
	for !hospitalQueue.IsEmpty() {
		patient, _ := hospitalQueue.Pop()
		fmt.Printf("    %s (Urgency: %d)\n", patient.Name, patient.UrgencyLevel)
	}

	// Example 6: Game leaderboard (max-heap by score)
	fmt.Println("\n6. Game Leaderboard (Max-Heap by Score):")
	scoreQueue := NewMaxQueue(ScoreByPoints)

	scoreQueue.Push(Score{PlayerName: "Alice", Points: 1500})
	scoreQueue.Push(Score{PlayerName: "Bob", Points: 2000})
	scoreQueue.Push(Score{PlayerName: "Charlie", Points: 1200})

	fmt.Println("  Leaderboard:")
	rank := 1
	for !scoreQueue.IsEmpty() {
		score, _ := scoreQueue.Pop()
		fmt.Printf("    %d. %s: %d points\n", rank, score.PlayerName, score.Points)
		rank++
	}

	// Example 7: Custom comparison with lambda-like function
	fmt.Println("\n7. Custom Comparison (Tasks by Name Length):")
	taskByNameLengthQueue := NewMinQueue(func(a, b Task) int {
		return IntCompare(len(a.Name), len(b.Name))
	})

	taskByNameLengthQueue.Push(Task{Name: "A", Priority: 1})
	taskByNameLengthQueue.Push(Task{Name: "Very long task name", Priority: 2})
	taskByNameLengthQueue.Push(Task{Name: "Medium", Priority: 3})

	for !taskByNameLengthQueue.IsEmpty() {
		task, _ := taskByNameLengthQueue.Pop()
		fmt.Printf("  Task: '%s' (Length: %d)\n", task.Name, len(task.Name))
	}

	// Example 8: Using ReverseCompare
	fmt.Println("\n8. Reverse Integer Comparison (Max-Heap using Min comparator):")
	reverseIntQueue := NewMinQueue(ReverseCompare(IntCompare))

	reverseIntQueue.Push(10)
	reverseIntQueue.Push(30)
	reverseIntQueue.Push(20)

	for !reverseIntQueue.IsEmpty() {
		val, _ := reverseIntQueue.Pop()
		fmt.Printf("  Popped: %d\n", val)
	}
}
