package stack

import (
	"fmt"
	"testing"
)

func TestBasicPushPop(t *testing.T) {
	s := NewStack[int]()

	// Test empty stack
	if !s.IsEmpty() {
		t.Error("Expected empty stack")
	}

	if s.Size() != 0 {
		t.Error("Expected size 0")
	}

	// Test push and pop
	s.Push(10)
	s.Push(20)
	s.Push(30)

	if s.Size() != 3 {
		t.Error("Expected size 3")
	}

	if s.IsEmpty() {
		t.Error("Expected non-empty stack")
	}

	// Should pop in LIFO order (30, 20, 10)
	val1, err := s.Pop()
	if err != nil || val1 != 30 {
		t.Errorf("Expected value 30, got %v with error %v", val1, err)
	}

	val2, err := s.Pop()
	if err != nil || val2 != 20 {
		t.Errorf("Expected value 20, got %v with error %v", val2, err)
	}

	val3, err := s.Pop()
	if err != nil || val3 != 10 {
		t.Errorf("Expected value 10, got %v with error %v", val3, err)
	}

	// Test empty pop
	_, err = s.Pop()
	if err == nil {
		t.Error("Expected error on empty pop")
	}

	if !s.IsEmpty() {
		t.Error("Expected empty stack after popping all items")
	}
}

func TestPeek(t *testing.T) {
	s := NewStack[string]()

	// Test empty peek
	_, err := s.Peek()
	if err == nil {
		t.Error("Expected error on empty peek")
	}

	// Test with single item
	s.Push("single")

	top, err := s.Peek()
	if err != nil || top != "single" {
		t.Errorf("Expected top 'single', got %v with error %v", top, err)
	}

	// Test with multiple items
	s.Push("second")
	s.Push("third")

	top, err = s.Peek()
	if err != nil || top != "third" {
		t.Errorf("Expected top 'third', got %v with error %v", top, err)
	}

	// Peek should not modify the stack
	if s.Size() != 3 {
		t.Error("Expected size 3 after peek operations")
	}
}

func TestCustomTypes(t *testing.T) {
	type Task struct {
		ID   int
		Name string
	}

	s := NewStack[Task]()

	t1 := Task{ID: 1, Name: "First task"}
	t2 := Task{ID: 2, Name: "Second task"}
	t3 := Task{ID: 3, Name: "Third task"}

	s.Push(t1)
	s.Push(t2)
	s.Push(t3)

	// Should pop in LIFO order
	task3, err := s.Pop()
	if err != nil || task3.ID != 3 {
		t.Errorf("Expected task 3, got %v with error %v", task3, err)
	}

	task2, err := s.Pop()
	if err != nil || task2.ID != 2 {
		t.Errorf("Expected task 2, got %v with error %v", task2, err)
	}

	task1, err := s.Pop()
	if err != nil || task1.ID != 1 {
		t.Errorf("Expected task 1, got %v with error %v", task1, err)
	}
}

func TestClear(t *testing.T) {
	s := NewStack[int]()

	s.Push(1)
	s.Push(2)
	s.Push(3)

	if s.Size() != 3 {
		t.Error("Expected size 3 before clear")
	}

	s.Clear()

	if !s.IsEmpty() || s.Size() != 0 {
		t.Error("Expected empty stack after clear")
	}

	// Should be able to use stack after clear
	s.Push(42)
	val, err := s.Pop()
	if err != nil || val != 42 {
		t.Error("Stack should be usable after clear")
	}
}

func TestLargeDataset(t *testing.T) {
	s := NewStack[int]()

	// Push 1000 items
	for i := 0; i < 1000; i++ {
		s.Push(i)
	}

	if s.Size() != 1000 {
		t.Errorf("Expected size 1000, got %d", s.Size())
	}

	// Pop all items and verify LIFO order
	for i := 999; i >= 0; i-- {
		val, err := s.Pop()
		if err != nil {
			t.Errorf("Unexpected error at position %d: %v", i, err)
		}
		if val != i {
			t.Errorf("Expected %d at position %d, got %d", i, i, val)
		}
	}

	if !s.IsEmpty() {
		t.Error("Expected empty stack after popping all items")
	}
}

func TestSingleItemOperations(t *testing.T) {
	s := NewStack[int]()

	// Test single item push/pop
	s.Push(42)

	if s.IsEmpty() {
		t.Error("Expected non-empty stack after push")
	}

	if s.Size() != 1 {
		t.Error("Expected size 1")
	}

	top, err := s.Peek()
	if err != nil || top != 42 {
		t.Error("Peek should return the single item")
	}

	val, err := s.Pop()
	if err != nil || val != 42 {
		t.Error("Pop should return the single item")
	}

	if !s.IsEmpty() {
		t.Error("Expected empty stack after popping single item")
	}
}

func TestParenthesesMatching(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{"()", true},
		{"(())", true},
		{"((()))", true},
		{"()()", true},
		{"(()())", true},
		{"(()", false},
		{"())", false},
		{")(", false},
		{"", true},
		{"((())", false},
		{"))(", false},
	}

	for _, tc := range testCases {
		result := isValidParenthesesTest(tc.input)
		if result != tc.expected {
			t.Errorf("For input '%s', expected %t, got %t", tc.input, tc.expected, result)
		}
	}
}

// Helper function for testing parentheses matching
func isValidParenthesesTest(s string) bool {
	stack := NewStack[rune]()

	for _, char := range s {
		switch char {
		case '(':
			stack.Push(char)
		case ')':
			if stack.IsEmpty() {
				return false
			}
			stack.Pop()
		}
	}

	return stack.IsEmpty()
}

func TestPostfixEvaluation(t *testing.T) {
	testCases := []struct {
		expression []string
		expected   int
	}{
		{[]string{"3", "4", "+"}, 7},                      // 3 + 4 = 7
		{[]string{"3", "4", "+", "2", "*"}, 14},           // (3 + 4) * 2 = 14
		{[]string{"5", "2", "-", "3", "*"}, 9},            // (5 - 2) * 3 = 9
		{[]string{"8", "2", "/"}, 4},                      // 8 / 2 = 4
		{[]string{"1", "2", "+", "3", "4", "+", "*"}, 21}, // (1 + 2) * (3 + 4) = 21
	}

	for _, tc := range testCases {
		result := evaluatePostfix(tc.expression)
		if result != tc.expected {
			t.Errorf("For expression %v, expected %d, got %d", tc.expression, tc.expected, result)
		}
	}
}

// Helper function for testing postfix evaluation
func evaluatePostfix(tokens []string) int {
	stack := NewStack[int]()

	for _, token := range tokens {
		switch token {
		case "+", "-", "*", "/":
			b, _ := stack.Pop()
			a, _ := stack.Pop()

			var result int
			switch token {
			case "+":
				result = a + b
			case "-":
				result = a - b
			case "*":
				result = a * b
			case "/":
				result = a / b
			}

			stack.Push(result)
		default:
			// Convert string to int
			var num int
			fmt.Sscanf(token, "%d", &num)
			stack.Push(num)
		}
	}

	result, _ := stack.Pop()
	return result
}

func TestEdgeCases(t *testing.T) {
	t.Run("Multiple push/pop cycles", func(t *testing.T) {
		s := NewStack[int]()

		// Multiple cycles of fill and empty
		for cycle := 0; cycle < 3; cycle++ {
			// Fill stack
			for i := 0; i < 5; i++ {
				s.Push(i + cycle*10)
			}

			// Empty stack
			for i := 4; i >= 0; i-- {
				val, err := s.Pop()
				if err != nil {
					t.Errorf("Cycle %d: unexpected error at pop %d: %v", cycle, i, err)
				}
				expected := i + cycle*10
				if val != expected {
					t.Errorf("Cycle %d: expected %d, got %d", cycle, expected, val)
				}
			}

			if !s.IsEmpty() {
				t.Errorf("Cycle %d: expected empty stack", cycle)
			}
		}
	})

	t.Run("Zero values", func(t *testing.T) {
		s := NewStack[int]()

		s.Push(0)
		s.Push(-1)
		s.Push(1)

		expected := []int{1, -1, 0} // LIFO order
		for i, exp := range expected {
			val, err := s.Pop()
			if err != nil || val != exp {
				t.Errorf("Expected %d at position %d, got %d", exp, i, val)
			}
		}
	})

	t.Run("Duplicate values", func(t *testing.T) {
		s := NewStack[string]()

		s.Push("same")
		s.Push("same")
		s.Push("same")

		for i := 0; i < 3; i++ {
			val, err := s.Pop()
			if err != nil || val != "same" {
				t.Errorf("Expected 'same' at position %d, got %v", i, val)
			}
		}
	})
}

// Benchmark tests
func BenchmarkPush(b *testing.B) {
	s := NewStack[int]()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Push(i)
	}
}

func BenchmarkPop(b *testing.B) {
	s := NewStack[int]()

	// Pre-populate
	for i := 0; i < b.N; i++ {
		s.Push(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Pop()
	}
}

func BenchmarkPeek(b *testing.B) {
	s := NewStack[int]()
	s.Push(42)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Peek()
	}
}

func BenchmarkPushPop(b *testing.B) {
	s := NewStack[int]()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Push(i)
		if i%2 == 0 { // Pop every other iteration
			s.Pop()
		}
	}
}

// Example tests for documentation
func ExampleNew() {
	s := NewStack[int]()

	s.Push(10)
	s.Push(20)
	s.Push(30)

	for !s.IsEmpty() {
		val, _ := s.Pop()
		println("Popped:", val)
	}
	// Output:
	// Popped: 30
	// Popped: 20
	// Popped: 10
}

func ExampleStack_Push() {
	s := NewStack[string]()

	s.Push("first")
	s.Push("second")
	s.Push("third")

	println("Stack size:", s.Size())
	// Output: Stack size: 3
}

func ExampleStack_Pop() {
	s := NewStack[string]()

	s.Push("A")
	s.Push("B")

	second, _ := s.Pop()
	first, _ := s.Pop()

	println("Second (last in):", second)
	println("First (first in):", first)
	// Output:
	// Second (last in): B
	// First (first in): A
}

func ExampleStack_Peek() {
	s := NewStack[int]()

	s.Push(100)
	s.Push(200)

	top, _ := s.Peek()
	println("Top item:", top)
	println("Stack size after Peek():", s.Size())
	// Output:
	// Top item: 200
	// Stack size after Peek(): 2
}
