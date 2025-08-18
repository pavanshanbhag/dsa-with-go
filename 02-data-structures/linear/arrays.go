package datastructures

import (
	"errors"
	"fmt"
)

// Dynamic Array (Slice-based) Implementation
// This demonstrates Go's slice mechanics and provides a foundation for understanding
// dynamic arrays with explicit capacity management

type DynamicArray struct {
	data     []interface{}
	size     int
	capacity int
}

// NewDynamicArray creates a new dynamic array with initial capacity
func NewDynamicArray(initialCapacity int) *DynamicArray {
	if initialCapacity <= 0 {
		initialCapacity = 1
	}
	return &DynamicArray{
		data:     make([]interface{}, initialCapacity),
		size:     0,
		capacity: initialCapacity,
	}
}

// Size returns the number of elements in the array
func (da *DynamicArray) Size() int {
	return da.size
}

// Capacity returns the current capacity of the array
func (da *DynamicArray) Capacity() int {
	return da.capacity
}

// IsEmpty checks if the array is empty
func (da *DynamicArray) IsEmpty() bool {
	return da.size == 0
}

// Get retrieves an element at the specified index
func (da *DynamicArray) Get(index int) (interface{}, error) {
	if index < 0 || index >= da.size {
		return nil, errors.New("index out of bounds")
	}
	return da.data[index], nil
}

// Set updates an element at the specified index
func (da *DynamicArray) Set(index int, value interface{}) error {
	if index < 0 || index >= da.size {
		return errors.New("index out of bounds")
	}
	da.data[index] = value
	return nil
}

// Add appends an element to the end of the array
func (da *DynamicArray) Add(value interface{}) {
	if da.size >= da.capacity {
		da.resize()
	}
	da.data[da.size] = value
	da.size++
}

// Insert adds an element at the specified index
func (da *DynamicArray) Insert(index int, value interface{}) error {
	if index < 0 || index > da.size {
		return errors.New("index out of bounds")
	}

	if da.size >= da.capacity {
		da.resize()
	}

	// Shift elements to the right
	for i := da.size; i > index; i-- {
		da.data[i] = da.data[i-1]
	}

	da.data[index] = value
	da.size++
	return nil
}

// Remove deletes an element at the specified index
func (da *DynamicArray) Remove(index int) (interface{}, error) {
	if index < 0 || index >= da.size {
		return nil, errors.New("index out of bounds")
	}

	removedValue := da.data[index]

	// Shift elements to the left
	for i := index; i < da.size-1; i++ {
		da.data[i] = da.data[i+1]
	}

	da.size--
	da.data[da.size] = nil // Clear reference to help GC

	// Shrink if necessary
	if da.size <= da.capacity/4 && da.capacity > 1 {
		da.shrink()
	}

	return removedValue, nil
}

// Pop removes and returns the last element
func (da *DynamicArray) Pop() (interface{}, error) {
	if da.size == 0 {
		return nil, errors.New("array is empty")
	}
	return da.Remove(da.size - 1)
}

// IndexOf finds the first occurrence of a value
func (da *DynamicArray) IndexOf(value interface{}) int {
	for i := 0; i < da.size; i++ {
		if da.data[i] == value {
			return i
		}
	}
	return -1
}

// Contains checks if the array contains a value
func (da *DynamicArray) Contains(value interface{}) bool {
	return da.IndexOf(value) != -1
}

// Clear removes all elements from the array
func (da *DynamicArray) Clear() {
	for i := 0; i < da.size; i++ {
		da.data[i] = nil
	}
	da.size = 0
}

// ToSlice returns a copy of the array as a slice
func (da *DynamicArray) ToSlice() []interface{} {
	result := make([]interface{}, da.size)
	copy(result, da.data[:da.size])
	return result
}

// String provides a string representation of the array
func (da *DynamicArray) String() string {
	if da.size == 0 {
		return "[]"
	}

	result := "["
	for i := 0; i < da.size; i++ {
		if i > 0 {
			result += ", "
		}
		result += fmt.Sprintf("%v", da.data[i])
	}
	result += "]"
	return result
}

// resize doubles the capacity of the array
func (da *DynamicArray) resize() {
	newCapacity := da.capacity * 2
	newData := make([]interface{}, newCapacity)
	copy(newData, da.data)
	da.data = newData
	da.capacity = newCapacity
}

// shrink halves the capacity of the array
func (da *DynamicArray) shrink() {
	newCapacity := da.capacity / 2
	newData := make([]interface{}, newCapacity)
	copy(newData, da.data[:da.size])
	da.data = newData
	da.capacity = newCapacity
}

// Iterator provides a way to iterate over the array
type DynamicArrayIterator struct {
	array *DynamicArray
	index int
}

// NewIterator creates a new iterator for the array
func (da *DynamicArray) NewIterator() *DynamicArrayIterator {
	return &DynamicArrayIterator{
		array: da,
		index: 0,
	}
}

// HasNext checks if there are more elements to iterate
func (iter *DynamicArrayIterator) HasNext() bool {
	return iter.index < iter.array.size
}

// Next returns the next element in the iteration
func (iter *DynamicArrayIterator) Next() (interface{}, error) {
	if !iter.HasNext() {
		return nil, errors.New("no more elements")
	}

	value := iter.array.data[iter.index]
	iter.index++
	return value, nil
}

// Reset resets the iterator to the beginning
func (iter *DynamicArrayIterator) Reset() {
	iter.index = 0
}

// Generic Dynamic Array for type safety (Go 1.18+)
type GenericDynamicArray[T any] struct {
	data     []T
	size     int
	capacity int
}

// NewGenericDynamicArray creates a new type-safe dynamic array
func NewGenericDynamicArray[T any](initialCapacity int) *GenericDynamicArray[T] {
	if initialCapacity <= 0 {
		initialCapacity = 1
	}
	return &GenericDynamicArray[T]{
		data:     make([]T, initialCapacity),
		size:     0,
		capacity: initialCapacity,
	}
}

// Add appends an element to the generic array
func (gda *GenericDynamicArray[T]) Add(value T) {
	if gda.size >= gda.capacity {
		gda.resize()
	}
	gda.data[gda.size] = value
	gda.size++
}

// Get retrieves an element at the specified index
func (gda *GenericDynamicArray[T]) Get(index int) (T, error) {
	var zero T
	if index < 0 || index >= gda.size {
		return zero, errors.New("index out of bounds")
	}
	return gda.data[index], nil
}

// Size returns the number of elements
func (gda *GenericDynamicArray[T]) Size() int {
	return gda.size
}

// ToSlice returns a copy as a slice
func (gda *GenericDynamicArray[T]) ToSlice() []T {
	result := make([]T, gda.size)
	copy(result, gda.data[:gda.size])
	return result
}

// resize doubles the capacity of the generic array
func (gda *GenericDynamicArray[T]) resize() {
	newCapacity := gda.capacity * 2
	newData := make([]T, newCapacity)
	copy(newData, gda.data)
	gda.data = newData
	gda.capacity = newCapacity
}
