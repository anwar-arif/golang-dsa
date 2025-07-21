package queue

import (
	"fmt"
	"testing"
)

func TestBasicEnqueueDequeue(t *testing.T) {
	q := NewQueue[int]()

	// Test empty queue
	if !q.IsEmpty() {
		t.Error("Expected empty queue")
	}

	if q.Size() != 0 {
		t.Error("Expected size 0")
	}

	// Test enqueue and dequeue
	q.Push(10)
	q.Push(20)
	q.Push(30)

	if q.Size() != 3 {
		t.Error("Expected size 3")
	}

	if q.IsEmpty() {
		t.Error("Expected non-empty queue")
	}

	// Should dequeue in FIFO order (10, 20, 30)
	val1, err := q.Pop()
	if err != nil || val1 != 10 {
		t.Errorf("Expected value 10, got %v with error %v", val1, err)
	}

	val2, err := q.Pop()
	if err != nil || val2 != 20 {
		t.Errorf("Expected value 20, got %v with error %v", val2, err)
	}

	val3, err := q.Pop()
	if err != nil || val3 != 30 {
		t.Errorf("Expected value 30, got %v with error %v", val3, err)
	}

	// Test empty dequeue
	_, err = q.Pop()
	if err == nil {
		t.Error("Expected error on empty dequeue")
	}

	if !q.IsEmpty() {
		t.Error("Expected empty queue after dequeuing all items")
	}
}

func TestFrontAndRear(t *testing.T) {
	q := NewQueue[string]()

	// Test empty queue
	_, err := q.Front()
	if err == nil {
		t.Error("Expected error on empty front")
	}

	_, err = q.Rear()
	if err == nil {
		t.Error("Expected error on empty rear")
	}

	// Test with single item
	q.Push("single")

	front, err := q.Front()
	if err != nil || front != "single" {
		t.Errorf("Expected front 'single', got %v with error %v", front, err)
	}

	rear, err := q.Rear()
	if err != nil || rear != "single" {
		t.Errorf("Expected rear 'single', got %v with error %v", rear, err)
	}

	// Test with multiple items
	q.Push("second")
	q.Push("third")

	front, err = q.Front()
	if err != nil || front != "single" {
		t.Errorf("Expected front 'single', got %v with error %v", front, err)
	}

	rear, err = q.Rear()
	if err != nil || rear != "third" {
		t.Errorf("Expected rear 'third', got %v with error %v", rear, err)
	}

	// Front and rear should not modify the queue
	if q.Size() != 3 {
		t.Error("Expected size 3 after front/rear operations")
	}
}

func TestCustomTypes(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	q := NewQueue[Person]()

	p1 := Person{Name: "Alice", Age: 30}
	p2 := Person{Name: "Bob", Age: 25}
	p3 := Person{Name: "Charlie", Age: 35}

	q.Push(p1)
	q.Push(p2)
	q.Push(p3)

	// Should dequeue in FIFO order
	person1, err := q.Pop()
	if err != nil || person1.Name != "Alice" {
		t.Errorf("Expected Alice, got %v with error %v", person1, err)
	}

	person2, err := q.Pop()
	if err != nil || person2.Name != "Bob" {
		t.Errorf("Expected Bob, got %v with error %v", person2, err)
	}

	person3, err := q.Pop()
	if err != nil || person3.Name != "Charlie" {
		t.Errorf("Expected Charlie, got %v with error %v", person3, err)
	}
}

func TestToSlice(t *testing.T) {
	q := NewQueue[int]()

	// Test empty queue
	slice := q.ToSlice()
	if len(slice) != 0 {
		t.Error("Expected empty slice for empty queue")
	}

	// Test with items
	q.Push(1)
	q.Push(2)
	q.Push(3)

	slice = q.ToSlice()
	expected := []int{1, 2, 3}

	if len(slice) != len(expected) {
		t.Errorf("Expected slice length %d, got %d", len(expected), len(slice))
	}

	for i, val := range expected {
		if slice[i] != val {
			t.Errorf("Expected slice[%d] = %d, got %d", i, val, slice[i])
		}
	}

	// ToSlice should not modify the queue
	if q.Size() != 3 {
		t.Error("Expected size 3 after ToSlice")
	}
}

func TestClear(t *testing.T) {
	q := NewQueue[int]()

	q.Push(1)
	q.Push(2)
	q.Push(3)

	if q.Size() != 3 {
		t.Error("Expected size 3 before clear")
	}

	q.Clear()

	if !q.IsEmpty() || q.Size() != 0 {
		t.Error("Expected empty queue after clear")
	}

	// Should be able to use queue after clear
	q.Push(42)
	val, err := q.Pop()
	if err != nil || val != 42 {
		t.Error("Queue should be usable after clear")
	}
}

func TestMixedOperations(t *testing.T) {
	q := NewQueue[string]()

	// Mix enqueue and dequeue operations
	q.Push("A")
	q.Push("B")

	val, _ := q.Pop()
	if val != "A" {
		t.Errorf("Expected 'A', got %s", val)
	}

	q.Push("C")
	q.Push("D")

	val, _ = q.Pop()
	if val != "B" {
		t.Errorf("Expected 'B', got %s", val)
	}

	val, _ = q.Pop()
	if val != "C" {
		t.Errorf("Expected 'C', got %s", val)
	}

	q.Push("E")

	expected := []string{"D", "E"}
	actual := q.ToSlice()

	for i, exp := range expected {
		if actual[i] != exp {
			t.Errorf("Expected %s at position %d, got %s", exp, i, actual[i])
		}
	}
}

func TestLargeDataset(t *testing.T) {
	q := NewQueue[int]()

	// Enqueue 1000 items
	for i := 0; i < 1000; i++ {
		q.Push(i)
	}

	if q.Size() != 1000 {
		t.Errorf("Expected size 1000, got %d", q.Size())
	}

	// Dequeue all items and verify FIFO order
	for i := 0; i < 1000; i++ {
		val, err := q.Pop()
		if err != nil {
			t.Errorf("Unexpected error at position %d: %v", i, err)
		}
		if val != i {
			t.Errorf("Expected %d at position %d, got %d", i, i, val)
		}
	}

	if !q.IsEmpty() {
		t.Error("Expected empty queue after dequeuing all items")
	}
}

func TestSingleItemOperations(t *testing.T) {
	q := NewQueue[int]()

	// Test single item enqueue/dequeue
	q.Push(42)

	if q.IsEmpty() {
		t.Error("Expected non-empty queue after enqueue")
	}

	if q.Size() != 1 {
		t.Error("Expected size 1")
	}

	front, err := q.Front()
	if err != nil || front != 42 {
		t.Error("Front should return the single item")
	}

	rear, err := q.Rear()
	if err != nil || rear != 42 {
		t.Error("Rear should return the single item")
	}

	val, err := q.Pop()
	if err != nil || val != 42 {
		t.Error("Dequeue should return the single item")
	}

	if !q.IsEmpty() {
		t.Error("Expected empty queue after dequeuing single item")
	}
}

func TestEdgeCases(t *testing.T) {
	t.Run("Multiple enqueue/dequeue cycles", func(t *testing.T) {
		q := NewQueue[int]()

		// Multiple cycles of fill and empty
		for cycle := 0; cycle < 3; cycle++ {
			// Fill queue
			for i := 0; i < 5; i++ {
				q.Push(i + cycle*10)
			}

			// Empty queue
			for i := 0; i < 5; i++ {
				val, err := q.Pop()
				if err != nil {
					t.Errorf("Cycle %d: unexpected error at dequeue %d: %v", cycle, i, err)
				}
				expected := i + cycle*10
				if val != expected {
					t.Errorf("Cycle %d: expected %d, got %d", cycle, expected, val)
				}
			}

			if !q.IsEmpty() {
				t.Errorf("Cycle %d: expected empty queue", cycle)
			}
		}
	})

	t.Run("Zero values", func(t *testing.T) {
		q := NewQueue[int]()

		q.Push(0)
		q.Push(-1)
		q.Push(1)

		expected := []int{0, -1, 1}
		for i, exp := range expected {
			val, err := q.Pop()
			if err != nil || val != exp {
				t.Errorf("Expected %d at position %d, got %d", exp, i, val)
			}
		}
	})

	t.Run("Duplicate values", func(t *testing.T) {
		q := NewQueue[string]()

		q.Push("same")
		q.Push("same")
		q.Push("same")

		for i := 0; i < 3; i++ {
			val, err := q.Pop()
			if err != nil || val != "same" {
				t.Errorf("Expected 'same' at position %d, got %v", i, val)
			}
		}
	})
}

// Benchmark tests
func BenchmarkEnqueue(b *testing.B) {
	q := NewQueue[int]()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		q.Push(i)
	}
}

func BenchmarkDequeue(b *testing.B) {
	q := NewQueue[int]()

	// Pre-populate
	for i := 0; i < b.N; i++ {
		q.Push(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		q.Pop()
	}
}

func BenchmarkFront(b *testing.B) {
	q := NewQueue[int]()
	q.Push(42)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		q.Front()
	}
}

func BenchmarkEnqueueDequeue(b *testing.B) {
	q := NewQueue[int]()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		q.Push(i)
		if i%2 == 0 { // Dequeue every other iteration
			q.Pop()
		}
	}
}

// Example tests for documentation
func ExampleNew() {
	q := NewQueue[int]()

	q.Push(10)
	q.Push(20)
	q.Push(30)

	for !q.IsEmpty() {
		val, _ := q.Pop()
		println("Dequeued:", val)
	}
	// Output:
	// Dequeued: 10
	// Dequeued: 20
	// Dequeued: 30
}

func ExampleQueue_Enqueue() {
	q := NewQueue[string]()

	q.Push("first")
	q.Push("second")
	q.Push("third")

	println("Queue size:", q.Size())
	// Output: Queue size: 3
}

func ExampleQueue_Dequeue() {
	q := NewQueue[string]()

	q.Push("A")
	q.Push("B")

	first, _ := q.Pop()
	second, _ := q.Pop()

	println("First:", first)
	println("Second:", second)
	// Output:
	// First: A
	// Second: B
}

func ExampleQueue_Front() {
	q := NewQueue[int]()

	q.Push(100)
	q.Push(200)

	front, _ := q.Front()
	println("Front item:", front)
	println("Queue size after Front():", q.Size())
	// Output:
	// Front item: 100
	// Queue size after Front(): 2
}

func ExampleQueue_ToSlice() {
	q := NewQueue[string]()

	q.Push("apple")
	q.Push("banana")
	q.Push("cherry")

	items := q.ToSlice()
	println("Queue contents:", items)
	// Output: Queue contents: [apple banana cherry]
}

func ExampleQueue_NewIterator() {
	q := NewQueue[int]()

	q.Push(1)
	q.Push(2)
	q.Push(3)

	for !q.IsEmpty() {
		fmt.Println("Item: ", q.front)
	}
	// Output:
	// Item: 1
	// Item: 2
	// Item: 3
}
