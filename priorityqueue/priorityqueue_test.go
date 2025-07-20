package priorityqueue

import (
	"testing"
)

func TestMinQueueInts(t *testing.T) {
	pq := NewMinQueue(IntCompare)

	// Test empty queue
	if !pq.IsEmpty() {
		t.Error("Expected empty queue")
	}

	if pq.Size() != 0 {
		t.Error("Expected size 0")
	}

	// Test push and pop
	pq.Push(30)
	pq.Push(10)
	pq.Push(20)

	if pq.Size() != 3 {
		t.Error("Expected size 3")
	}

	// Should pop in ascending order (10, 20, 30)
	val1, err := pq.Pop()
	if err != nil || val1 != 10 {
		t.Errorf("Expected value 10, got %v with error %v", val1, err)
	}

	val2, err := pq.Pop()
	if err != nil || val2 != 20 {
		t.Errorf("Expected value 20, got %v with error %v", val2, err)
	}

	val3, err := pq.Pop()
	if err != nil || val3 != 30 {
		t.Errorf("Expected value 30, got %v with error %v", val3, err)
	}

	// Test empty pop
	_, err = pq.Pop()
	if err == nil {
		t.Error("Expected error on empty pop")
	}
}

func TestMaxQueueStrings(t *testing.T) {
	pq := NewMaxQueue(StringCompare)

	pq.Push("apple")
	pq.Push("zebra")
	pq.Push("banana")

	// Should pop in descending order (zebra, banana, apple)
	val1, err := pq.Pop()
	if err != nil || val1 != "zebra" {
		t.Errorf("Expected value 'zebra', got %v with error %v", val1, err)
	}

	val2, err := pq.Pop()
	if err != nil || val2 != "banana" {
		t.Errorf("Expected value 'banana', got %v with error %v", val2, err)
	}

	val3, err := pq.Pop()
	if err != nil || val3 != "apple" {
		t.Errorf("Expected value 'apple', got %v with error %v", val3, err)
	}
}

func TestTaskQueue(t *testing.T) {
	pq := NewMinQueue(TaskByPriority)

	pq.Push(Task{ID: 1, Name: "Task1", Priority: 3})
	pq.Push(Task{ID: 2, Name: "Task2", Priority: 1})
	pq.Push(Task{ID: 3, Name: "Task3", Priority: 2})

	// Should pop in priority order (1, 2, 3)
	task1, err := pq.Pop()
	if err != nil || task1.Priority != 1 || task1.ID != 2 {
		t.Errorf("Expected priority 1, ID 2, got priority %d, ID %d", task1.Priority, task1.ID)
	}

	task2, err := pq.Pop()
	if err != nil || task2.Priority != 2 || task2.ID != 3 {
		t.Errorf("Expected priority 2, ID 3, got priority %d, ID %d", task2.Priority, task2.ID)
	}

	task3, err := pq.Pop()
	if err != nil || task3.Priority != 3 || task3.ID != 1 {
		t.Errorf("Expected priority 3, ID 1, got priority %d, ID %d", task3.Priority, task3.ID)
	}
}

func TestPeek(t *testing.T) {
	pq := NewMinQueue(IntCompare)

	// Test empty peek
	_, err := pq.Peek()
	if err == nil {
		t.Error("Expected error on empty peek")
	}

	pq.Push(10)
	pq.Push(5)

	// Peek should return highest priority without removing
	val, err := pq.Peek()
	if err != nil || val != 5 {
		t.Errorf("Expected value 5, got %v with error %v", val, err)
	}

	// Size should remain same
	if pq.Size() != 2 {
		t.Error("Expected size 2 after peek")
	}

	// Pop should return same item
	val, err = pq.Pop()
	if err != nil || val != 5 {
		t.Errorf("Expected value 5, got %v with error %v", val, err)
	}
}

func TestCustomComparison(t *testing.T) {
	// Custom comparison: compare by string length
	lengthCompare := func(a, b string) int {
		return IntCompare(len(a), len(b))
	}

	pq := NewMinQueue(lengthCompare)

	pq.Push("hello")  // length 5
	pq.Push("hi")     // length 2
	pq.Push("world!") // length 6

	// Should pop in length order: "hi", "hello", "world!"
	val1, _ := pq.Pop()
	if val1 != "hi" {
		t.Errorf("Expected 'hi' (shortest), got %v", val1)
	}

	val2, _ := pq.Pop()
	if val2 != "hello" {
		t.Errorf("Expected 'hello' (medium), got %v", val2)
	}

	val3, _ := pq.Pop()
	if val3 != "world!" {
		t.Errorf("Expected 'world!' (longest), got %v", val3)
	}
}

func TestReverseCompare(t *testing.T) {
	// Using ReverseCompare to make a max-heap with min-heap constructor
	pq := NewMinQueue(ReverseCompare(IntCompare))

	pq.Push(10)
	pq.Push(30)
	pq.Push(20)

	// Should pop in descending order (30, 20, 10)
	val1, _ := pq.Pop()
	if val1 != 30 {
		t.Errorf("Expected 30, got %v", val1)
	}

	val2, _ := pq.Pop()
	if val2 != 20 {
		t.Errorf("Expected 20, got %v", val2)
	}

	val3, _ := pq.Pop()
	if val3 != 10 {
		t.Errorf("Expected 10, got %v", val3)
	}
}

func TestFloatComparison(t *testing.T) {
	pq := NewMinQueue(Float64Compare)

	pq.Push(3.14)
	pq.Push(2.71)
	pq.Push(1.41)

	val1, _ := pq.Pop()
	if val1 != 1.41 {
		t.Errorf("Expected 1.41, got %v", val1)
	}

	val2, _ := pq.Pop()
	if val2 != 2.71 {
		t.Errorf("Expected 2.71, got %v", val2)
	}

	val3, _ := pq.Pop()
	if val3 != 3.14 {
		t.Errorf("Expected 3.14, got %v", val3)
	}
}

func TestDijkstraScenario(t *testing.T) {
	pq := NewMinQueue(NodeByDistance)

	// Simulate Dijkstra's algorithm
	pq.Push(Node{ID: 0, Distance: 0})
	pq.Push(Node{ID: 1, Distance: 4})
	pq.Push(Node{ID: 2, Distance: 2})
	pq.Push(Node{ID: 3, Distance: 6})
	pq.Push(Node{ID: 4, Distance: 1})

	// Should process in distance order: 0, 1, 2, 4, 6
	expectedDistances := []int{0, 1, 2, 4, 6}
	expectedIDs := []int{0, 4, 2, 1, 3}

	for i, expectedDist := range expectedDistances {
		node, err := pq.Pop()
		if err != nil {
			t.Errorf("Unexpected error at position %d: %v", i, err)
		}
		if node.Distance != expectedDist {
			t.Errorf("Expected distance %d at position %d, got %d", expectedDist, i, node.Distance)
		}
		if node.ID != expectedIDs[i] {
			t.Errorf("Expected ID %d at position %d, got %d", expectedIDs[i], i, node.ID)
		}
	}
}

func TestHospitalScenario(t *testing.T) {
	pq := NewMaxQueue(PatientByUrgency)

	// Emergency room scenario
	pq.Push(Patient{Name: "John", Age: 25, UrgencyLevel: 3})
	pq.Push(Patient{Name: "Sarah", Age: 45, UrgencyLevel: 8})
	pq.Push(Patient{Name: "Mike", Age: 30, UrgencyLevel: 1})
	pq.Push(Patient{Name: "Anna", Age: 60, UrgencyLevel: 5})

	// Should process in urgency order: 8, 5, 3, 1
	expectedUrgencies := []int{8, 5, 3, 1}
	expectedNames := []string{"Sarah", "Anna", "John", "Mike"}

	for i, expectedUrgency := range expectedUrgencies {
		patient, err := pq.Pop()
		if err != nil {
			t.Errorf("Unexpected error at position %d: %v", i, err)
		}
		if patient.UrgencyLevel != expectedUrgency {
			t.Errorf("Expected urgency %d at position %d, got %d", expectedUrgency, i, patient.UrgencyLevel)
		}
		if patient.Name != expectedNames[i] {
			t.Errorf("Expected name %s at position %d, got %s", expectedNames[i], i, patient.Name)
		}
	}
}

func TestMultiCriteriaComparison(t *testing.T) {
	// Test jobs with priority and duration comparison
	type Job struct {
		ID       int
		Priority int
		Duration int
	}

	jobCompare := func(a, b Job) int {
		// First compare by priority (lower = higher precedence)
		if cmp := IntCompare(a.Priority, b.Priority); cmp != 0 {
			return cmp
		}
		// If priorities are equal, compare by duration (shorter first)
		return IntCompare(a.Duration, b.Duration)
	}

	pq := NewMinQueue(jobCompare)

	pq.Push(Job{ID: 1, Priority: 2, Duration: 10})
	pq.Push(Job{ID: 2, Priority: 1, Duration: 20})
	pq.Push(Job{ID: 3, Priority: 1, Duration: 5}) // Same priority as job 2, but shorter
	pq.Push(Job{ID: 4, Priority: 3, Duration: 1})

	// Expected order: Job3 (p:1,d:5), Job2 (p:1,d:20), Job1 (p:2,d:10), Job4 (p:3,d:1)
	expectedIDs := []int{3, 2, 1, 4}

	for i, expectedID := range expectedIDs {
		job, err := pq.Pop()
		if err != nil {
			t.Errorf("Unexpected error at position %d: %v", i, err)
		}
		if job.ID != expectedID {
			t.Errorf("Expected job ID %d at position %d, got %d", expectedID, i, job.ID)
		}
	}
}

func TestRemove(t *testing.T) {
	pq := NewMinQueue(IntCompare)

	pq.Push(10)
	pq.Push(20)
	pq.Push(30)

	items := pq.ToSlice()

	// Remove an item (this is a bit tricky to test since heap order isn't guaranteed)
	if len(items) >= 2 {
		originalSize := pq.Size()
		pq.Remove(items[1])

		if pq.Size() != originalSize-1 {
			t.Errorf("Expected size %d after remove, got %d", originalSize-1, pq.Size())
		}

		// Pop all remaining items
		remaining := make([]int, 0)
		for !pq.IsEmpty() {
			val, _ := pq.Pop()
			remaining = append(remaining, val)
		}

		// Should have exactly originalSize-1 items
		if len(remaining) != originalSize-1 {
			t.Errorf("Expected %d remaining items, got %d", originalSize-1, len(remaining))
		}
	}
}

func TestClear(t *testing.T) {
	pq := NewMinQueue(IntCompare)

	pq.Push(10)
	pq.Push(20)
	pq.Push(30)

	if pq.Size() != 3 {
		t.Error("Expected size 3 before clear")
	}

	pq.Clear()

	if !pq.IsEmpty() || pq.Size() != 0 {
		t.Error("Expected empty queue after clear")
	}

	// Should be able to use queue after clear
	pq.Push(42)
	val, err := pq.Pop()
	if err != nil || val != 42 {
		t.Error("Queue should be usable after clear")
	}
}

func TestUpdateItem(t *testing.T) {
	// Test with mutable tasks
	type MutableTask struct {
		ID       int
		Name     string
		Priority int
	}

	taskCompare := func(a, b *MutableTask) int {
		return IntCompare(a.Priority, b.Priority)
	}

	pq := NewMinQueue(taskCompare)

	task1 := &MutableTask{ID: 1, Name: "Task1", Priority: 3}
	task2 := &MutableTask{ID: 2, Name: "Task2", Priority: 2}
	task3 := &MutableTask{ID: 3, Name: "Task3", Priority: 1}

	pq.Push(task1)
	pq.Push(task2)
	pq.Push(task3)

	// Get items from heap
	items := pq.ToSlice()

	// Find task1 and update its priority
	for _, item := range items {
		if item.Value.ID == 1 {
			item.Value.Priority = 0 // Make it highest priority
			pq.UpdateItem(item)
			break
		}
	}

	// Pop should return the updated task
	popped, err := pq.Pop()
	if err != nil {
		t.Error("Expected successful pop")
	}

	if popped.ID != 1 {
		t.Errorf("Expected updated task with ID 1, got ID %d", popped.ID)
	}

	if popped.Priority != 0 {
		t.Errorf("Expected updated priority 0, got %d", popped.Priority)
	}
}

func TestLargeDataset(t *testing.T) {
	pq := NewMinQueue(IntCompare)

	// Push 1000 random-ish numbers
	testData := []int{500, 100, 800, 200, 700, 300, 600, 400, 900, 50}
	for i := 0; i < 100; i++ {
		for _, val := range testData {
			pq.Push(val + i*1000)
		}
	}

	if pq.Size() != 1000 {
		t.Errorf("Expected size 1000, got %d", pq.Size())
	}

	// Pop all items and verify they come out in sorted order
	var prev int
	var count int
	for !pq.IsEmpty() {
		val, err := pq.Pop()
		if err != nil {
			t.Errorf("Unexpected error popping item %d: %v", count, err)
		}

		if count > 0 && val < prev {
			t.Errorf("Items not in sorted order: prev=%d, current=%d at position %d", prev, val, count)
		}

		prev = val
		count++
	}

	if count != 1000 {
		t.Errorf("Expected to pop 1000 items, got %d", count)
	}
}

func TestEdgeCases(t *testing.T) {
	t.Run("Single item", func(t *testing.T) {
		pq := NewMinQueue(IntCompare)
		pq.Push(42)

		val, err := pq.Peek()
		if err != nil || val != 42 {
			t.Error("Failed to peek single item")
		}

		val, err = pq.Pop()
		if err != nil || val != 42 {
			t.Error("Failed to pop single item")
		}

		if !pq.IsEmpty() {
			t.Error("Queue should be empty after popping single item")
		}
	})

	t.Run("Duplicate values", func(t *testing.T) {
		pq := NewMinQueue(IntCompare)

		pq.Push(5)
		pq.Push(5)
		pq.Push(5)

		for i := 0; i < 3; i++ {
			val, err := pq.Pop()
			if err != nil || val != 5 {
				t.Errorf("Expected value 5 at position %d, got %v", i, val)
			}
		}
	})

	t.Run("Zero values", func(t *testing.T) {
		pq := NewMinQueue(IntCompare)

		pq.Push(0)
		pq.Push(-1)
		pq.Push(1)

		expected := []int{-1, 0, 1}
		for i, exp := range expected {
			val, err := pq.Pop()
			if err != nil || val != exp {
				t.Errorf("Expected %d at position %d, got %v", exp, i, val)
			}
		}
	})
}

// Benchmark tests
func BenchmarkPushMinQueue(b *testing.B) {
	pq := NewMinQueue(IntCompare)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pq.Push(i)
	}
}

func BenchmarkPopMinQueue(b *testing.B) {
	pq := NewMinQueue(IntCompare)

	// Pre-populate
	for i := 0; i < b.N; i++ {
		pq.Push(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pq.Pop()
	}
}

func BenchmarkPeekMinQueue(b *testing.B) {
	pq := NewMinQueue(IntCompare)
	pq.Push(42)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pq.Peek()
	}
}

func BenchmarkStringQueue(b *testing.B) {
	pq := NewMinQueue(StringCompare)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pq.Push("test string")
	}
}

// Example test for documentation
func ExampleNewMinQueue() {
	pq := NewMinQueue(IntCompare)

	pq.Push(30)
	pq.Push(10)
	pq.Push(20)

	for !pq.IsEmpty() {
		val, _ := pq.Pop()
		println("Value:", val)
	}
	// Output:
	// Value: 10
	// Value: 20
	// Value: 30
}

func ExampleNewMaxQueue() {
	pq := NewMaxQueue(StringCompare)

	pq.Push("apple")
	pq.Push("zebra")
	pq.Push("banana")

	for !pq.IsEmpty() {
		val, _ := pq.Pop()
		println("Value:", val)
	}
	// Output:
	// Value: zebra
	// Value: banana
	// Value: apple
}

func ExampleCustomComparison() {
	// Custom comparison by string length
	lengthCompare := func(a, b string) int {
		return IntCompare(len(a), len(b))
	}

	pq := NewMinQueue(lengthCompare)

	pq.Push("hello")
	pq.Push("hi")
	pq.Push("world!")

	for !pq.IsEmpty() {
		val, _ := pq.Pop()
		println("Value:", val, "Length:", len(val))
	}
	// Output:
	// Value: hi Length: 2
	// Value: hello Length: 5
	// Value: world! Length: 6
}
