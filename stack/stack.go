package stack

import (
	"fmt"
)

// Node represents a node in the stack
type Node[T any] struct {
	Value T
	Next  *Node[T]
}

// Stack represents a LIFO stack
type Stack[T any] struct {
	top  *Node[T] // Points to the top element (push/pop from here)
	size int
}

// NewStack creates a new empty stack
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		top:  nil,
		size: 0,
	}
}

// Push adds an item to the top of the stack
func (s *Stack[T]) Push(value T) {
	newNode := &Node[T]{
		Value: value,
		Next:  s.top, // Point to the previous top
	}

	s.top = newNode
	s.size++
}

// Pop removes and returns the item from the top of the stack
func (s *Stack[T]) Pop() (T, error) {
	var zero T

	if s.IsEmpty() {
		return zero, fmt.Errorf("stack is empty")
	}

	value := s.top.Value
	s.top = s.top.Next
	s.size--

	return value, nil
}

// Peek returns the top item without removing it
func (s *Stack[T]) Peek() (T, error) {
	var zero T

	if s.IsEmpty() {
		return zero, fmt.Errorf("stack is empty")
	}

	return s.top.Value, nil
}

// IsEmpty returns true if the stack is empty
func (s *Stack[T]) IsEmpty() bool {
	return s.top == nil
}

// Size returns the number of items in the stack
func (s *Stack[T]) Size() int {
	return s.size
}

// Clear removes all items from the stack
func (s *Stack[T]) Clear() {
	s.top = nil
	s.size = 0
}

// Example usage and demonstrations
func ExampleUsage() {
	fmt.Println("=== Generic Stack Examples ===\n")

	// Example 1: Basic integer stack
	fmt.Println("1. Basic Integer Stack (LIFO):")
	intStack := NewStack[int]()

	intStack.Push(10)
	intStack.Push(20)
	intStack.Push(30)

	fmt.Printf("Stack size: %d\n", intStack.Size())
	fmt.Printf("Top item: ")
	if top, err := intStack.Peek(); err == nil {
		fmt.Printf("%d\n", top)
	}

	fmt.Println("Popping items:")
	for !intStack.IsEmpty() {
		val, _ := intStack.Pop()
		fmt.Printf("  Popped: %d\n", val)
	}

	// Example 2: String stack
	fmt.Println("\n2. String Stack:")
	stringStack := NewStack[string]()

	stringStack.Push("first")
	stringStack.Push("second")
	stringStack.Push("third")

	for !stringStack.IsEmpty() {
		val, _ := stringStack.Pop()
		fmt.Printf("  Popped: %s\n", val)
	}

	// Example 3: Function call stack simulation
	fmt.Println("\n3. Function Call Stack Simulation:")

	type FunctionCall struct {
		Name string
		Args []string
	}

	callStack := NewStack[FunctionCall]()

	callStack.Push(FunctionCall{Name: "main", Args: []string{}})
	callStack.Push(FunctionCall{Name: "processData", Args: []string{"data.txt"}})
	callStack.Push(FunctionCall{Name: "validateInput", Args: []string{"input"}})

	fmt.Println("Function call stack:")
	for !callStack.IsEmpty() {
		call, _ := callStack.Pop()
		fmt.Printf("  Returning from: %s(%v)\n", call.Name, call.Args)
	}

	// Example 4: Stack operations
	fmt.Println("\n4. Stack Operations:")
	opStack := NewStack[string]()

	opStack.Push("A")
	opStack.Push("B")
	opStack.Push("C")

	// Example 6: Expression evaluation (postfix)
	fmt.Println("\n6. Postfix Expression Evaluation:")

	type Token struct {
		Value string
		Type  string // "number" or "operator"
	}

	evalStack := NewStack[int]()

	// Example: "3 4 + 2 *" = (3 + 4) * 2 = 14
	tokens := []Token{
		{Value: "3", Type: "number"},
		{Value: "4", Type: "number"},
		{Value: "+", Type: "operator"},
		{Value: "2", Type: "number"},
		{Value: "*", Type: "operator"},
	}

	fmt.Print("Evaluating postfix expression: ")
	for _, token := range tokens {
		fmt.Printf("%s ", token.Value)
	}
	fmt.Println()

	for _, token := range tokens {
		if token.Type == "number" {
			// Convert string to int (simplified)
			var num int
			fmt.Sscanf(token.Value, "%d", &num)
			evalStack.Push(num)
		} else {
			// Pop two operands
			b, _ := evalStack.Pop()
			a, _ := evalStack.Pop()
			var result int

			switch token.Value {
			case "+":
				result = a + b
			case "-":
				result = a - b
			case "*":
				result = a * b
			case "/":
				result = a / b
			}

			evalStack.Push(result)
		}
	}

	finalResult, _ := evalStack.Pop()
	fmt.Printf("Final result: %d\n", finalResult)

	// Example 7: Parentheses matching
	fmt.Println("\n7. Parentheses Matching:")

	expressions := []string{
		"((()))",
		"((())",
		"()()())",
		"(()())",
		"(()",
	}

	for _, expr := range expressions {
		valid := isValidParentheses(expr)
		fmt.Printf("  '%s' -> %t\n", expr, valid)
	}

	// Example 8: Error handling
	fmt.Println("\n8. Error Handling:")
	emptyStack := NewStack[int]()

	_, err := emptyStack.Pop()
	if err != nil {
		fmt.Printf("Pop error: %v\n", err)
	}

	_, err = emptyStack.Peek()
	if err != nil {
		fmt.Printf("Peek error: %v\n", err)
	}
}

// Helper function for parentheses matching
func isValidParentheses(s string) bool {
	stack := NewStack[rune]()

	pairs := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, char := range s {
		switch char {
		case '(', '[', '{':
			stack.Push(char)
		case ')', ']', '}':
			if stack.IsEmpty() {
				return false
			}

			top, _ := stack.Pop()
			if top != pairs[char] {
				return false
			}
		}
	}

	return stack.IsEmpty()
}