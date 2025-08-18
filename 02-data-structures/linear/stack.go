package datastructures

import (
	"errors"
	"fmt"
	"sync"
)

// Stack interface defines the basic stack operations
type Stack interface {
	Push(value interface{})
	Pop() (interface{}, error)
	Peek() (interface{}, error)
	Size() int
	IsEmpty() bool
	Clear()
}

// ArrayStack implements a stack using a dynamic array (slice)
type ArrayStack struct {
	data []interface{}
	top  int
}

// NewArrayStack creates a new array-based stack
func NewArrayStack() *ArrayStack {
	return &ArrayStack{
		data: make([]interface{}, 0),
		top:  -1,
	}
}

// Push adds an element to the top of the stack
func (s *ArrayStack) Push(value interface{}) {
	s.data = append(s.data, value)
	s.top++
}

// Pop removes and returns the top element
func (s *ArrayStack) Pop() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("stack is empty")
	}

	value := s.data[s.top]
	s.data = s.data[:s.top]
	s.top--
	return value, nil
}

// Peek returns the top element without removing it
func (s *ArrayStack) Peek() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("stack is empty")
	}
	return s.data[s.top], nil
}

// Size returns the number of elements in the stack
func (s *ArrayStack) Size() int {
	return s.top + 1
}

// IsEmpty checks if the stack is empty
func (s *ArrayStack) IsEmpty() bool {
	return s.top == -1
}

// Clear removes all elements from the stack
func (s *ArrayStack) Clear() {
	s.data = s.data[:0]
	s.top = -1
}

// Linked Stack Node
type StackNode struct {
	data interface{}
	next *StackNode
}

// LinkedStack implements a stack using linked nodes
type LinkedStack struct {
	top  *StackNode
	size int
}

// NewLinkedStack creates a new linked stack
func NewLinkedStack() *LinkedStack {
	return &LinkedStack{
		top:  nil,
		size: 0,
	}
}

// Push adds an element to the top of the stack
func (s *LinkedStack) Push(value interface{}) {
	newNode := &StackNode{
		data: value,
		next: s.top,
	}
	s.top = newNode
	s.size++
}

// Pop removes and returns the top element
func (s *LinkedStack) Pop() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("stack is empty")
	}

	value := s.top.data
	s.top = s.top.next
	s.size--
	return value, nil
}

// Peek returns the top element without removing it
func (s *LinkedStack) Peek() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("stack is empty")
	}
	return s.top.data, nil
}

// Size returns the number of elements in the stack
func (s *LinkedStack) Size() int {
	return s.size
}

// IsEmpty checks if the stack is empty
func (s *LinkedStack) IsEmpty() bool {
	return s.top == nil
}

// Clear removes all elements from the stack
func (s *LinkedStack) Clear() {
	s.top = nil
	s.size = 0
}

// Thread-Safe Stack
type SafeStack struct {
	stack Stack
	mutex sync.RWMutex
}

// NewSafeStack creates a new thread-safe stack
func NewSafeStack(underlying Stack) *SafeStack {
	return &SafeStack{
		stack: underlying,
	}
}

// Push adds an element to the stack (thread-safe)
func (s *SafeStack) Push(value interface{}) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.stack.Push(value)
}

// Pop removes and returns the top element (thread-safe)
func (s *SafeStack) Pop() (interface{}, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.stack.Pop()
}

// Peek returns the top element without removing it (thread-safe)
func (s *SafeStack) Peek() (interface{}, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return s.stack.Peek()
}

// Size returns the number of elements (thread-safe)
func (s *SafeStack) Size() int {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return s.stack.Size()
}

// IsEmpty checks if the stack is empty (thread-safe)
func (s *SafeStack) IsEmpty() bool {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return s.stack.IsEmpty()
}

// Clear removes all elements (thread-safe)
func (s *SafeStack) Clear() {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.stack.Clear()
}

// Generic Stack for type safety
type GenericStack[T any] struct {
	data []T
}

// NewGenericStack creates a new generic stack
func NewGenericStack[T any]() *GenericStack[T] {
	return &GenericStack[T]{
		data: make([]T, 0),
	}
}

// Push adds an element to the stack
func (s *GenericStack[T]) Push(value T) {
	s.data = append(s.data, value)
}

// Pop removes and returns the top element
func (s *GenericStack[T]) Pop() (T, error) {
	var zero T
	if len(s.data) == 0 {
		return zero, errors.New("stack is empty")
	}

	index := len(s.data) - 1
	value := s.data[index]
	s.data = s.data[:index]
	return value, nil
}

// Peek returns the top element without removing it
func (s *GenericStack[T]) Peek() (T, error) {
	var zero T
	if len(s.data) == 0 {
		return zero, errors.New("stack is empty")
	}
	return s.data[len(s.data)-1], nil
}

// Size returns the number of elements
func (s *GenericStack[T]) Size() int {
	return len(s.data)
}

// IsEmpty checks if the stack is empty
func (s *GenericStack[T]) IsEmpty() bool {
	return len(s.data) == 0
}

// Stack Applications

// BalancedParentheses checks if parentheses are balanced using a stack
func BalancedParentheses(expression string) bool {
	stack := NewGenericStack[rune]()
	pairs := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, char := range expression {
		switch char {
		case '(', '{', '[':
			stack.Push(char)
		case ')', '}', ']':
			if stack.IsEmpty() {
				return false
			}

			top, err := stack.Pop()
			if err != nil || top != pairs[char] {
				return false
			}
		}
	}

	return stack.IsEmpty()
}

// EvaluatePostfix evaluates a postfix expression using a stack
func EvaluatePostfix(expression []string) (int, error) {
	stack := NewGenericStack[int]()

	for _, token := range expression {
		switch token {
		case "+":
			b, err1 := stack.Pop()
			a, err2 := stack.Pop()
			if err1 != nil || err2 != nil {
				return 0, errors.New("invalid expression")
			}
			stack.Push(a + b)
		case "-":
			b, err1 := stack.Pop()
			a, err2 := stack.Pop()
			if err1 != nil || err2 != nil {
				return 0, errors.New("invalid expression")
			}
			stack.Push(a - b)
		case "*":
			b, err1 := stack.Pop()
			a, err2 := stack.Pop()
			if err1 != nil || err2 != nil {
				return 0, errors.New("invalid expression")
			}
			stack.Push(a * b)
		case "/":
			b, err1 := stack.Pop()
			a, err2 := stack.Pop()
			if err1 != nil || err2 != nil || b == 0 {
				return 0, errors.New("invalid expression or division by zero")
			}
			stack.Push(a / b)
		default:
			// Assume it's a number
			var num int
			if _, err := fmt.Sscanf(token, "%d", &num); err != nil {
				return 0, errors.New("invalid number: " + token)
			}
			stack.Push(num)
		}
	}

	if stack.Size() != 1 {
		return 0, errors.New("invalid expression")
	}

	return stack.Pop()
}
