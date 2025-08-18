package datastructures

import (
	"errors"
	"sync"
)

// Queue interface defines the basic queue operations
type Queue interface {
	Enqueue(value interface{}) error
	Dequeue() (interface{}, error)
	Front() (interface{}, error)
	Size() int
	IsEmpty() bool
	Clear()
}

// ArrayQueue implements a queue using a circular buffer
type ArrayQueue struct {
	data     []interface{}
	front    int
	rear     int
	size     int
	capacity int
}

// NewArrayQueue creates a new array-based queue with specified capacity
func NewArrayQueue(capacity int) *ArrayQueue {
	if capacity <= 0 {
		capacity = 10
	}
	return &ArrayQueue{
		data:     make([]interface{}, capacity),
		front:    0,
		rear:     0,
		size:     0,
		capacity: capacity,
	}
}

// Enqueue adds an element to the rear of the queue
func (q *ArrayQueue) Enqueue(value interface{}) error {
	if q.size >= q.capacity {
		return errors.New("queue is full")
	}

	q.data[q.rear] = value
	q.rear = (q.rear + 1) % q.capacity
	q.size++
	return nil
}

// Dequeue removes and returns the front element
func (q *ArrayQueue) Dequeue() (interface{}, error) {
	if q.IsEmpty() {
		return nil, errors.New("queue is empty")
	}

	value := q.data[q.front]
	q.data[q.front] = nil // Help GC
	q.front = (q.front + 1) % q.capacity
	q.size--
	return value, nil
}

// Front returns the front element without removing it
func (q *ArrayQueue) Front() (interface{}, error) {
	if q.IsEmpty() {
		return nil, errors.New("queue is empty")
	}
	return q.data[q.front], nil
}

// Size returns the number of elements in the queue
func (q *ArrayQueue) Size() int {
	return q.size
}

// IsEmpty checks if the queue is empty
func (q *ArrayQueue) IsEmpty() bool {
	return q.size == 0
}

// IsFull checks if the queue is full
func (q *ArrayQueue) IsFull() bool {
	return q.size >= q.capacity
}

// Clear removes all elements from the queue
func (q *ArrayQueue) Clear() {
	for i := 0; i < q.capacity; i++ {
		q.data[i] = nil
	}
	q.front = 0
	q.rear = 0
	q.size = 0
}

// Dynamic Array Queue (resizable)
type DynamicArrayQueue struct {
	data []interface{}
}

// NewDynamicArrayQueue creates a new resizable queue
func NewDynamicArrayQueue() *DynamicArrayQueue {
	return &DynamicArrayQueue{
		data: make([]interface{}, 0),
	}
}

// Enqueue adds an element to the rear
func (q *DynamicArrayQueue) Enqueue(value interface{}) error {
	q.data = append(q.data, value)
	return nil
}

// Dequeue removes and returns the front element
func (q *DynamicArrayQueue) Dequeue() (interface{}, error) {
	if q.IsEmpty() {
		return nil, errors.New("queue is empty")
	}

	value := q.data[0]
	q.data = q.data[1:]
	return value, nil
}

// Front returns the front element without removing it
func (q *DynamicArrayQueue) Front() (interface{}, error) {
	if q.IsEmpty() {
		return nil, errors.New("queue is empty")
	}
	return q.data[0], nil
}

// Size returns the number of elements
func (q *DynamicArrayQueue) Size() int {
	return len(q.data)
}

// IsEmpty checks if the queue is empty
func (q *DynamicArrayQueue) IsEmpty() bool {
	return len(q.data) == 0
}

// Clear removes all elements
func (q *DynamicArrayQueue) Clear() {
	q.data = q.data[:0]
}

// Linked Queue Node
type QueueNode struct {
	data interface{}
	next *QueueNode
}

// LinkedQueue implements a queue using linked nodes
type LinkedQueue struct {
	front *QueueNode
	rear  *QueueNode
	size  int
}

// NewLinkedQueue creates a new linked queue
func NewLinkedQueue() *LinkedQueue {
	return &LinkedQueue{
		front: nil,
		rear:  nil,
		size:  0,
	}
}

// Enqueue adds an element to the rear
func (q *LinkedQueue) Enqueue(value interface{}) error {
	newNode := &QueueNode{
		data: value,
		next: nil,
	}

	if q.IsEmpty() {
		q.front = newNode
		q.rear = newNode
	} else {
		q.rear.next = newNode
		q.rear = newNode
	}

	q.size++
	return nil
}

// Dequeue removes and returns the front element
func (q *LinkedQueue) Dequeue() (interface{}, error) {
	if q.IsEmpty() {
		return nil, errors.New("queue is empty")
	}

	value := q.front.data
	q.front = q.front.next

	if q.front == nil {
		q.rear = nil
	}

	q.size--
	return value, nil
}

// Front returns the front element without removing it
func (q *LinkedQueue) Front() (interface{}, error) {
	if q.IsEmpty() {
		return nil, errors.New("queue is empty")
	}
	return q.front.data, nil
}

// Size returns the number of elements
func (q *LinkedQueue) Size() int {
	return q.size
}

// IsEmpty checks if the queue is empty
func (q *LinkedQueue) IsEmpty() bool {
	return q.front == nil
}

// Clear removes all elements
func (q *LinkedQueue) Clear() {
	q.front = nil
	q.rear = nil
	q.size = 0
}

// Priority Queue implementation using heap
type PriorityItem struct {
	Value    interface{}
	Priority int
}

type PriorityQueue struct {
	items []PriorityItem
}

// NewPriorityQueue creates a new priority queue
func NewPriorityQueue() *PriorityQueue {
	return &PriorityQueue{
		items: make([]PriorityItem, 0),
	}
}

// Enqueue adds an element with priority (higher number = higher priority)
func (pq *PriorityQueue) Enqueue(value interface{}, priority int) error {
	item := PriorityItem{Value: value, Priority: priority}
	pq.items = append(pq.items, item)
	pq.heapifyUp(len(pq.items) - 1)
	return nil
}

// Dequeue removes and returns the highest priority element
func (pq *PriorityQueue) Dequeue() (interface{}, error) {
	if pq.IsEmpty() {
		return nil, errors.New("priority queue is empty")
	}

	// Get the root (highest priority)
	value := pq.items[0].Value

	// Move last element to root and remove it
	lastIndex := len(pq.items) - 1
	pq.items[0] = pq.items[lastIndex]
	pq.items = pq.items[:lastIndex]

	// Restore heap property
	if len(pq.items) > 0 {
		pq.heapifyDown(0)
	}

	return value, nil
}

// Front returns the highest priority element without removing it
func (pq *PriorityQueue) Front() (interface{}, error) {
	if pq.IsEmpty() {
		return nil, errors.New("priority queue is empty")
	}
	return pq.items[0].Value, nil
}

// Size returns the number of elements
func (pq *PriorityQueue) Size() int {
	return len(pq.items)
}

// IsEmpty checks if the queue is empty
func (pq *PriorityQueue) IsEmpty() bool {
	return len(pq.items) == 0
}

// Clear removes all elements
func (pq *PriorityQueue) Clear() {
	pq.items = pq.items[:0]
}

// heapifyUp maintains heap property upward
func (pq *PriorityQueue) heapifyUp(index int) {
	parentIndex := (index - 1) / 2

	if index > 0 && pq.items[index].Priority > pq.items[parentIndex].Priority {
		pq.items[index], pq.items[parentIndex] = pq.items[parentIndex], pq.items[index]
		pq.heapifyUp(parentIndex)
	}
}

// heapifyDown maintains heap property downward
func (pq *PriorityQueue) heapifyDown(index int) {
	leftChild := 2*index + 1
	rightChild := 2*index + 2
	largest := index

	if leftChild < len(pq.items) && pq.items[leftChild].Priority > pq.items[largest].Priority {
		largest = leftChild
	}

	if rightChild < len(pq.items) && pq.items[rightChild].Priority > pq.items[largest].Priority {
		largest = rightChild
	}

	if largest != index {
		pq.items[index], pq.items[largest] = pq.items[largest], pq.items[index]
		pq.heapifyDown(largest)
	}
}

// Thread-Safe Queue
type SafeQueue struct {
	queue Queue
	mutex sync.RWMutex
}

// NewSafeQueue creates a new thread-safe queue
func NewSafeQueue(underlying Queue) *SafeQueue {
	return &SafeQueue{
		queue: underlying,
	}
}

// Enqueue adds an element (thread-safe)
func (sq *SafeQueue) Enqueue(value interface{}) error {
	sq.mutex.Lock()
	defer sq.mutex.Unlock()
	return sq.queue.Enqueue(value)
}

// Dequeue removes and returns front element (thread-safe)
func (sq *SafeQueue) Dequeue() (interface{}, error) {
	sq.mutex.Lock()
	defer sq.mutex.Unlock()
	return sq.queue.Dequeue()
}

// Front returns front element (thread-safe)
func (sq *SafeQueue) Front() (interface{}, error) {
	sq.mutex.RLock()
	defer sq.mutex.RUnlock()
	return sq.queue.Front()
}

// Size returns number of elements (thread-safe)
func (sq *SafeQueue) Size() int {
	sq.mutex.RLock()
	defer sq.mutex.RUnlock()
	return sq.queue.Size()
}

// IsEmpty checks if empty (thread-safe)
func (sq *SafeQueue) IsEmpty() bool {
	sq.mutex.RLock()
	defer sq.mutex.RUnlock()
	return sq.queue.IsEmpty()
}

// Clear removes all elements (thread-safe)
func (sq *SafeQueue) Clear() {
	sq.mutex.Lock()
	defer sq.mutex.Unlock()
	sq.queue.Clear()
}

// Channel-based Queue (Go's idiomatic approach)
type ChannelQueue struct {
	ch       chan interface{}
	capacity int
}

// NewChannelQueue creates a new channel-based queue
func NewChannelQueue(capacity int) *ChannelQueue {
	return &ChannelQueue{
		ch:       make(chan interface{}, capacity),
		capacity: capacity,
	}
}

// Enqueue adds an element (non-blocking if capacity available)
func (cq *ChannelQueue) Enqueue(value interface{}) error {
	select {
	case cq.ch <- value:
		return nil
	default:
		return errors.New("queue is full")
	}
}

// EnqueueBlocking adds an element (blocks if full)
func (cq *ChannelQueue) EnqueueBlocking(value interface{}) error {
	cq.ch <- value
	return nil
}

// Dequeue removes and returns front element (non-blocking)
func (cq *ChannelQueue) Dequeue() (interface{}, error) {
	select {
	case value := <-cq.ch:
		return value, nil
	default:
		return nil, errors.New("queue is empty")
	}
}

// DequeueBlocking removes and returns front element (blocks if empty)
func (cq *ChannelQueue) DequeueBlocking() (interface{}, error) {
	value := <-cq.ch
	return value, nil
}

// Size returns approximate number of elements
func (cq *ChannelQueue) Size() int {
	return len(cq.ch)
}

// IsEmpty checks if the queue is empty
func (cq *ChannelQueue) IsEmpty() bool {
	return len(cq.ch) == 0
}

// IsFull checks if the queue is full
func (cq *ChannelQueue) IsFull() bool {
	return len(cq.ch) >= cq.capacity
}

// Close closes the underlying channel
func (cq *ChannelQueue) Close() {
	close(cq.ch)
}

// Generic Queue for type safety
type GenericQueue[T any] struct {
	data []T
}

// NewGenericQueue creates a new generic queue
func NewGenericQueue[T any]() *GenericQueue[T] {
	return &GenericQueue[T]{
		data: make([]T, 0),
	}
}

// Enqueue adds an element to the rear
func (gq *GenericQueue[T]) Enqueue(value T) {
	gq.data = append(gq.data, value)
}

// Dequeue removes and returns the front element
func (gq *GenericQueue[T]) Dequeue() (T, error) {
	var zero T
	if len(gq.data) == 0 {
		return zero, errors.New("queue is empty")
	}

	value := gq.data[0]
	gq.data = gq.data[1:]
	return value, nil
}

// Front returns the front element without removing it
func (gq *GenericQueue[T]) Front() (T, error) {
	var zero T
	if len(gq.data) == 0 {
		return zero, errors.New("queue is empty")
	}
	return gq.data[0], nil
}

// Size returns the number of elements
func (gq *GenericQueue[T]) Size() int {
	return len(gq.data)
}

// IsEmpty checks if the queue is empty
func (gq *GenericQueue[T]) IsEmpty() bool {
	return len(gq.data) == 0
}

// Clear removes all elements
func (gq *GenericQueue[T]) Clear() {
	gq.data = gq.data[:0]
}
