package datastructures

import (
	"errors"
	"fmt"
)

// Singly Linked List Node
type ListNode struct {
	Data interface{}
	Next *ListNode
}

// SinglyLinkedList implementation
type SinglyLinkedList struct {
	head *ListNode
	tail *ListNode
	size int
}

// NewSinglyLinkedList creates a new singly linked list
func NewSinglyLinkedList() *SinglyLinkedList {
	return &SinglyLinkedList{
		head: nil,
		tail: nil,
		size: 0,
	}
}

// Size returns the number of elements
func (sll *SinglyLinkedList) Size() int {
	return sll.size
}

// IsEmpty checks if the list is empty
func (sll *SinglyLinkedList) IsEmpty() bool {
	return sll.size == 0
}

// PrependFirst adds an element to the beginning
func (sll *SinglyLinkedList) PrependFirst(data interface{}) {
	newNode := &ListNode{Data: data, Next: sll.head}
	sll.head = newNode

	if sll.tail == nil {
		sll.tail = newNode
	}

	sll.size++
}

// AppendLast adds an element to the end
func (sll *SinglyLinkedList) AppendLast(data interface{}) {
	newNode := &ListNode{Data: data, Next: nil}

	if sll.tail == nil {
		sll.head = newNode
		sll.tail = newNode
	} else {
		sll.tail.Next = newNode
		sll.tail = newNode
	}

	sll.size++
}

// Get retrieves element at specified index
func (sll *SinglyLinkedList) Get(index int) (interface{}, error) {
	if index < 0 || index >= sll.size {
		return nil, errors.New("index out of bounds")
	}

	current := sll.head
	for i := 0; i < index; i++ {
		current = current.Next
	}

	return current.Data, nil
}

// Set updates element at specified index
func (sll *SinglyLinkedList) Set(index int, data interface{}) error {
	if index < 0 || index >= sll.size {
		return errors.New("index out of bounds")
	}

	current := sll.head
	for i := 0; i < index; i++ {
		current = current.Next
	}

	current.Data = data
	return nil
}

// Insert adds element at specified index
func (sll *SinglyLinkedList) Insert(index int, data interface{}) error {
	if index < 0 || index > sll.size {
		return errors.New("index out of bounds")
	}

	if index == 0 {
		sll.PrependFirst(data)
		return nil
	}

	if index == sll.size {
		sll.AppendLast(data)
		return nil
	}

	newNode := &ListNode{Data: data}
	current := sll.head

	for i := 0; i < index-1; i++ {
		current = current.Next
	}

	newNode.Next = current.Next
	current.Next = newNode
	sll.size++

	return nil
}

// RemoveFirst removes the first element
func (sll *SinglyLinkedList) RemoveFirst() (interface{}, error) {
	if sll.IsEmpty() {
		return nil, errors.New("list is empty")
	}

	data := sll.head.Data
	sll.head = sll.head.Next

	if sll.head == nil {
		sll.tail = nil
	}

	sll.size--
	return data, nil
}

// RemoveLast removes the last element
func (sll *SinglyLinkedList) RemoveLast() (interface{}, error) {
	if sll.IsEmpty() {
		return nil, errors.New("list is empty")
	}

	if sll.size == 1 {
		return sll.RemoveFirst()
	}

	current := sll.head
	for current.Next != sll.tail {
		current = current.Next
	}

	data := sll.tail.Data
	current.Next = nil
	sll.tail = current
	sll.size--

	return data, nil
}

// Remove element at specified index
func (sll *SinglyLinkedList) Remove(index int) (interface{}, error) {
	if index < 0 || index >= sll.size {
		return nil, errors.New("index out of bounds")
	}

	if index == 0 {
		return sll.RemoveFirst()
	}

	if index == sll.size-1 {
		return sll.RemoveLast()
	}

	current := sll.head
	for i := 0; i < index-1; i++ {
		current = current.Next
	}

	nodeToRemove := current.Next
	current.Next = nodeToRemove.Next
	sll.size--

	return nodeToRemove.Data, nil
}

// IndexOf finds the first occurrence of data
func (sll *SinglyLinkedList) IndexOf(data interface{}) int {
	current := sll.head
	index := 0

	for current != nil {
		if current.Data == data {
			return index
		}
		current = current.Next
		index++
	}

	return -1
}

// Contains checks if the list contains the data
func (sll *SinglyLinkedList) Contains(data interface{}) bool {
	return sll.IndexOf(data) != -1
}

// Clear removes all elements
func (sll *SinglyLinkedList) Clear() {
	sll.head = nil
	sll.tail = nil
	sll.size = 0
}

// ToSlice converts the list to a slice
func (sll *SinglyLinkedList) ToSlice() []interface{} {
	result := make([]interface{}, sll.size)
	current := sll.head
	index := 0

	for current != nil {
		result[index] = current.Data
		current = current.Next
		index++
	}

	return result
}

// String returns string representation
func (sll *SinglyLinkedList) String() string {
	if sll.IsEmpty() {
		return "[]"
	}

	result := "["
	current := sll.head

	for current != nil {
		result += fmt.Sprintf("%v", current.Data)
		if current.Next != nil {
			result += " -> "
		}
		current = current.Next
	}

	result += "]"
	return result
}

// Doubly Linked List Node
type DoublyListNode struct {
	Data interface{}
	Next *DoublyListNode
	Prev *DoublyListNode
}

// DoublyLinkedList implementation
type DoublyLinkedList struct {
	head *DoublyListNode
	tail *DoublyListNode
	size int
}

// NewDoublyLinkedList creates a new doubly linked list
func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{
		head: nil,
		tail: nil,
		size: 0,
	}
}

// Size returns the number of elements
func (dll *DoublyLinkedList) Size() int {
	return dll.size
}

// IsEmpty checks if the list is empty
func (dll *DoublyLinkedList) IsEmpty() bool {
	return dll.size == 0
}

// PrependFirst adds element to the beginning
func (dll *DoublyLinkedList) PrependFirst(data interface{}) {
	newNode := &DoublyListNode{Data: data, Next: dll.head, Prev: nil}

	if dll.head != nil {
		dll.head.Prev = newNode
	}

	dll.head = newNode

	if dll.tail == nil {
		dll.tail = newNode
	}

	dll.size++
}

// AppendLast adds element to the end
func (dll *DoublyLinkedList) AppendLast(data interface{}) {
	newNode := &DoublyListNode{Data: data, Next: nil, Prev: dll.tail}

	if dll.tail != nil {
		dll.tail.Next = newNode
	}

	dll.tail = newNode

	if dll.head == nil {
		dll.head = newNode
	}

	dll.size++
}

// Get retrieves element at specified index
func (dll *DoublyLinkedList) Get(index int) (interface{}, error) {
	if index < 0 || index >= dll.size {
		return nil, errors.New("index out of bounds")
	}

	var current *DoublyListNode

	// Optimize by choosing direction based on index
	if index < dll.size/2 {
		// Start from head
		current = dll.head
		for i := 0; i < index; i++ {
			current = current.Next
		}
	} else {
		// Start from tail
		current = dll.tail
		for i := dll.size - 1; i > index; i-- {
			current = current.Prev
		}
	}

	return current.Data, nil
}

// RemoveFirst removes the first element
func (dll *DoublyLinkedList) RemoveFirst() (interface{}, error) {
	if dll.IsEmpty() {
		return nil, errors.New("list is empty")
	}

	data := dll.head.Data
	dll.head = dll.head.Next

	if dll.head != nil {
		dll.head.Prev = nil
	} else {
		dll.tail = nil
	}

	dll.size--
	return data, nil
}

// RemoveLast removes the last element
func (dll *DoublyLinkedList) RemoveLast() (interface{}, error) {
	if dll.IsEmpty() {
		return nil, errors.New("list is empty")
	}

	data := dll.tail.Data
	dll.tail = dll.tail.Prev

	if dll.tail != nil {
		dll.tail.Next = nil
	} else {
		dll.head = nil
	}

	dll.size--
	return data, nil
}

// String returns string representation
func (dll *DoublyLinkedList) String() string {
	if dll.IsEmpty() {
		return "[]"
	}

	result := "["
	current := dll.head

	for current != nil {
		result += fmt.Sprintf("%v", current.Data)
		if current.Next != nil {
			result += " <-> "
		}
		current = current.Next
	}

	result += "]"
	return result
}

// Circular Linked List
type CircularLinkedList struct {
	tail *ListNode
	size int
}

// NewCircularLinkedList creates a new circular linked list
func NewCircularLinkedList() *CircularLinkedList {
	return &CircularLinkedList{
		tail: nil,
		size: 0,
	}
}

// Size returns the number of elements
func (cll *CircularLinkedList) Size() int {
	return cll.size
}

// IsEmpty checks if the list is empty
func (cll *CircularLinkedList) IsEmpty() bool {
	return cll.size == 0
}

// Add element to the list
func (cll *CircularLinkedList) Add(data interface{}) {
	newNode := &ListNode{Data: data}

	if cll.tail == nil {
		newNode.Next = newNode
		cll.tail = newNode
	} else {
		newNode.Next = cll.tail.Next
		cll.tail.Next = newNode
		cll.tail = newNode
	}

	cll.size++
}

// Remove removes the first occurrence of data
func (cll *CircularLinkedList) Remove(data interface{}) bool {
	if cll.IsEmpty() {
		return false
	}

	if cll.size == 1 {
		if cll.tail.Data == data {
			cll.tail = nil
			cll.size--
			return true
		}
		return false
	}

	current := cll.tail.Next // Start from head
	prev := cll.tail

	for i := 0; i < cll.size; i++ {
		if current.Data == data {
			if current == cll.tail {
				cll.tail = prev
			}
			prev.Next = current.Next
			cll.size--
			return true
		}
		prev = current
		current = current.Next
	}

	return false
}

// String returns string representation
func (cll *CircularLinkedList) String() string {
	if cll.IsEmpty() {
		return "[]"
	}

	result := "["
	current := cll.tail.Next // Start from head

	for i := 0; i < cll.size; i++ {
		result += fmt.Sprintf("%v", current.Data)
		if i < cll.size-1 {
			result += " -> "
		}
		current = current.Next
	}

	result += " -> (circular)]"
	return result
}

// Generic Linked List for type safety
type GenericLinkedList[T any] struct {
	head *GenericListNode[T]
	tail *GenericListNode[T]
	size int
}

type GenericListNode[T any] struct {
	Data T
	Next *GenericListNode[T]
}

// NewGenericLinkedList creates a new generic linked list
func NewGenericLinkedList[T any]() *GenericLinkedList[T] {
	return &GenericLinkedList[T]{
		head: nil,
		tail: nil,
		size: 0,
	}
}

// Append adds element to the end
func (gll *GenericLinkedList[T]) Append(data T) {
	newNode := &GenericListNode[T]{Data: data, Next: nil}

	if gll.tail == nil {
		gll.head = newNode
		gll.tail = newNode
	} else {
		gll.tail.Next = newNode
		gll.tail = newNode
	}

	gll.size++
}

// Get retrieves element at index
func (gll *GenericLinkedList[T]) Get(index int) (T, error) {
	var zero T
	if index < 0 || index >= gll.size {
		return zero, errors.New("index out of bounds")
	}

	current := gll.head
	for i := 0; i < index; i++ {
		current = current.Next
	}

	return current.Data, nil
}

// Size returns the number of elements
func (gll *GenericLinkedList[T]) Size() int {
	return gll.size
}

// IsEmpty checks if the list is empty
func (gll *GenericLinkedList[T]) IsEmpty() bool {
	return gll.size == 0
}

// ToSlice converts to slice
func (gll *GenericLinkedList[T]) ToSlice() []T {
	result := make([]T, gll.size)
	current := gll.head
	index := 0

	for current != nil {
		result[index] = current.Data
		current = current.Next
		index++
	}

	return result
}

// Linked List Applications

// Reverse a singly linked list
func (sll *SinglyLinkedList) Reverse() {
	if sll.size <= 1 {
		return
	}

	var prev *ListNode
	current := sll.head
	sll.tail = sll.head

	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}

	sll.head = prev
}

// Detect cycle in linked list (Floyd's algorithm)
func (sll *SinglyLinkedList) HasCycle() bool {
	if sll.head == nil || sll.head.Next == nil {
		return false
	}

	slow := sll.head
	fast := sll.head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true
		}
	}

	return false
}

// Find middle node
func (sll *SinglyLinkedList) FindMiddle() *ListNode {
	if sll.head == nil {
		return nil
	}

	slow := sll.head
	fast := sll.head

	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}
