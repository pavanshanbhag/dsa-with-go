package datastructures

import (
	"errors"
	"fmt"
)

// Binary Tree Node
type TreeNode struct {
	Data  interface{}
	Left  *TreeNode
	Right *TreeNode
}

// Binary Tree implementation
type BinaryTree struct {
	root *TreeNode
	size int
}

// NewBinaryTree creates a new binary tree
func NewBinaryTree() *BinaryTree {
	return &BinaryTree{
		root: nil,
		size: 0,
	}
}

// Size returns the number of nodes in the tree
func (bt *BinaryTree) Size() int {
	return bt.size
}

// IsEmpty checks if the tree is empty
func (bt *BinaryTree) IsEmpty() bool {
	return bt.root == nil
}

// Height returns the height of the tree
func (bt *BinaryTree) Height() int {
	return bt.height(bt.root)
}

func (bt *BinaryTree) height(node *TreeNode) int {
	if node == nil {
		return -1
	}

	leftHeight := bt.height(node.Left)
	rightHeight := bt.height(node.Right)

	if leftHeight > rightHeight {
		return leftHeight + 1
	}
	return rightHeight + 1
}

// Tree Traversal Methods

// InorderTraversal performs inorder traversal (Left, Root, Right)
func (bt *BinaryTree) InorderTraversal() []interface{} {
	var result []interface{}
	bt.inorder(bt.root, &result)
	return result
}

func (bt *BinaryTree) inorder(node *TreeNode, result *[]interface{}) {
	if node != nil {
		bt.inorder(node.Left, result)
		*result = append(*result, node.Data)
		bt.inorder(node.Right, result)
	}
}

// PreorderTraversal performs preorder traversal (Root, Left, Right)
func (bt *BinaryTree) PreorderTraversal() []interface{} {
	var result []interface{}
	bt.preorder(bt.root, &result)
	return result
}

func (bt *BinaryTree) preorder(node *TreeNode, result *[]interface{}) {
	if node != nil {
		*result = append(*result, node.Data)
		bt.preorder(node.Left, result)
		bt.preorder(node.Right, result)
	}
}

// PostorderTraversal performs postorder traversal (Left, Right, Root)
func (bt *BinaryTree) PostorderTraversal() []interface{} {
	var result []interface{}
	bt.postorder(bt.root, &result)
	return result
}

func (bt *BinaryTree) postorder(node *TreeNode, result *[]interface{}) {
	if node != nil {
		bt.postorder(node.Left, result)
		bt.postorder(node.Right, result)
		*result = append(*result, node.Data)
	}
}

// LevelOrderTraversal traverses the tree level by level (BFS)
func (bt *BinaryTree) LevelOrderTraversal() []interface{} {
	var result []interface{}
	if bt.root == nil {
		return result
	}

	// Simple queue implementation for tree traversal
	queue := []*TreeNode{bt.root}

	for len(queue) > 0 {
		// Dequeue front element
		node := queue[0]
		queue = queue[1:]

		// Process current node
		result = append(result, node.Data)

		// Enqueue children
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	return result
}

// Binary Search Tree (extends Binary Tree)
type BinarySearchTree struct {
	*BinaryTree
}

// NewBinarySearchTree creates a new BST
func NewBinarySearchTree() *BinarySearchTree {
	return &BinarySearchTree{
		BinaryTree: NewBinaryTree(),
	}
}

// Insert adds a value to the BST
func (bst *BinarySearchTree) Insert(data interface{}) error {
	var err error
	bst.root, err = bst.insertNode(bst.root, data)
	if err == nil {
		bst.size++
	}
	return err
}

func (bst *BinarySearchTree) insertNode(node *TreeNode, data interface{}) (*TreeNode, error) {
	if node == nil {
		return &TreeNode{Data: data}, nil
	}

	// For simplicity, assume data is comparable (int)
	nodeVal, ok1 := node.Data.(int)
	dataVal, ok2 := data.(int)

	if !ok1 || !ok2 {
		return node, errors.New("data must be comparable (int)")
	}

	if dataVal < nodeVal {
		var err error
		node.Left, err = bst.insertNode(node.Left, data)
		return node, err
	} else if dataVal > nodeVal {
		var err error
		node.Right, err = bst.insertNode(node.Right, data)
		return node, err
	}

	// Equal values - no insertion
	return node, nil
}

// Search finds a value in the BST
func (bst *BinarySearchTree) Search(data interface{}) bool {
	return bst.searchNode(bst.root, data)
}

func (bst *BinarySearchTree) searchNode(node *TreeNode, data interface{}) bool {
	if node == nil {
		return false
	}

	nodeVal, ok1 := node.Data.(int)
	dataVal, ok2 := data.(int)

	if !ok1 || !ok2 {
		return false
	}

	if dataVal == nodeVal {
		return true
	} else if dataVal < nodeVal {
		return bst.searchNode(node.Left, data)
	}

	return bst.searchNode(node.Right, data)
}

// FindMin finds the minimum value in the BST
func (bst *BinarySearchTree) FindMin() (interface{}, error) {
	if bst.root == nil {
		return nil, errors.New("tree is empty")
	}

	node := bst.findMinNode(bst.root)
	return node.Data, nil
}

func (bst *BinarySearchTree) findMinNode(node *TreeNode) *TreeNode {
	for node.Left != nil {
		node = node.Left
	}
	return node
}

// FindMax finds the maximum value in the BST
func (bst *BinarySearchTree) FindMax() (interface{}, error) {
	if bst.root == nil {
		return nil, errors.New("tree is empty")
	}

	node := bst.findMaxNode(bst.root)
	return node.Data, nil
}

func (bst *BinarySearchTree) findMaxNode(node *TreeNode) *TreeNode {
	for node.Right != nil {
		node = node.Right
	}
	return node
}

// Delete removes a value from the BST
func (bst *BinarySearchTree) Delete(data interface{}) error {
	var err error
	var deleted bool
	bst.root, deleted, err = bst.deleteNode(bst.root, data)
	if deleted {
		bst.size--
	}
	return err
}

func (bst *BinarySearchTree) deleteNode(node *TreeNode, data interface{}) (*TreeNode, bool, error) {
	if node == nil {
		return nil, false, nil
	}

	nodeVal, ok1 := node.Data.(int)
	dataVal, ok2 := data.(int)

	if !ok1 || !ok2 {
		return node, false, errors.New("data must be comparable (int)")
	}

	if dataVal < nodeVal {
		var deleted bool
		var err error
		node.Left, deleted, err = bst.deleteNode(node.Left, data)
		return node, deleted, err
	} else if dataVal > nodeVal {
		var deleted bool
		var err error
		node.Right, deleted, err = bst.deleteNode(node.Right, data)
		return node, deleted, err
	}

	// Node to be deleted found
	if node.Left == nil {
		return node.Right, true, nil
	} else if node.Right == nil {
		return node.Left, true, nil
	}

	// Node with two children
	successor := bst.findMinNode(node.Right)
	node.Data = successor.Data
	node.Right, _, _ = bst.deleteNode(node.Right, successor.Data)

	return node, true, nil
}

// IsValidBST checks if the tree maintains BST property
func (bst *BinarySearchTree) IsValidBST() bool {
	return bst.isValidBSTHelper(bst.root, nil, nil)
}

func (bst *BinarySearchTree) isValidBSTHelper(node *TreeNode, min, max *int) bool {
	if node == nil {
		return true
	}

	nodeVal, ok := node.Data.(int)
	if !ok {
		return false
	}

	if (min != nil && nodeVal <= *min) || (max != nil && nodeVal >= *max) {
		return false
	}

	return bst.isValidBSTHelper(node.Left, min, &nodeVal) &&
		bst.isValidBSTHelper(node.Right, &nodeVal, max)
}

// Heap implementation (Min Heap)
type MinHeap struct {
	items []int
}

// NewMinHeap creates a new min heap
func NewMinHeap() *MinHeap {
	return &MinHeap{
		items: make([]int, 0),
	}
}

// Size returns the number of elements
func (h *MinHeap) Size() int {
	return len(h.items)
}

// IsEmpty checks if heap is empty
func (h *MinHeap) IsEmpty() bool {
	return len(h.items) == 0
}

// Insert adds an element to the heap
func (h *MinHeap) Insert(value int) {
	h.items = append(h.items, value)
	h.heapifyUp(len(h.items) - 1)
}

// ExtractMin removes and returns the minimum element
func (h *MinHeap) ExtractMin() (int, error) {
	if h.IsEmpty() {
		return 0, errors.New("heap is empty")
	}

	min := h.items[0]
	lastIndex := len(h.items) - 1
	h.items[0] = h.items[lastIndex]
	h.items = h.items[:lastIndex]

	if len(h.items) > 0 {
		h.heapifyDown(0)
	}

	return min, nil
}

// Peek returns the minimum element without removing it
func (h *MinHeap) Peek() (int, error) {
	if h.IsEmpty() {
		return 0, errors.New("heap is empty")
	}
	return h.items[0], nil
}

// heapifyUp maintains heap property upward
func (h *MinHeap) heapifyUp(index int) {
	parentIndex := (index - 1) / 2

	if index > 0 && h.items[index] < h.items[parentIndex] {
		h.items[index], h.items[parentIndex] = h.items[parentIndex], h.items[index]
		h.heapifyUp(parentIndex)
	}
}

// heapifyDown maintains heap property downward
func (h *MinHeap) heapifyDown(index int) {
	leftChild := 2*index + 1
	rightChild := 2*index + 2
	smallest := index

	if leftChild < len(h.items) && h.items[leftChild] < h.items[smallest] {
		smallest = leftChild
	}

	if rightChild < len(h.items) && h.items[rightChild] < h.items[smallest] {
		smallest = rightChild
	}

	if smallest != index {
		h.items[index], h.items[smallest] = h.items[smallest], h.items[index]
		h.heapifyDown(smallest)
	}
}

// Build heap from slice
func (h *MinHeap) BuildHeap(arr []int) {
	h.items = make([]int, len(arr))
	copy(h.items, arr)

	// Start from last non-leaf node and heapify down
	for i := len(h.items)/2 - 1; i >= 0; i-- {
		h.heapifyDown(i)
	}
}

// HeapSort sorts an array using heap sort algorithm
func HeapSort(arr []int) []int {
	result := make([]int, len(arr))
	heap := NewMinHeap()

	// Insert all elements into heap
	for _, val := range arr {
		heap.Insert(val)
	}

	// Extract all elements (will be in sorted order)
	for i := 0; i < len(arr); i++ {
		min, _ := heap.ExtractMin()
		result[i] = min
	}

	return result
}

// String representation for trees
func (bt *BinaryTree) String() string {
	if bt.root == nil {
		return "Empty tree"
	}

	return bt.stringHelper(bt.root, "", true)
}

func (bt *BinaryTree) stringHelper(node *TreeNode, prefix string, isLast bool) string {
	if node == nil {
		return ""
	}

	result := prefix
	if isLast {
		result += "└── "
	} else {
		result += "├── "
	}
	result += fmt.Sprintf("%v\n", node.Data)

	children := []*TreeNode{node.Left, node.Right}
	for i, child := range children {
		if child != nil {
			childPrefix := prefix
			if isLast {
				childPrefix += "    "
			} else {
				childPrefix += "│   "
			}
			result += bt.stringHelper(child, childPrefix, i == 1 || children[1] == nil)
		}
	}

	return result
}

// Generic Binary Search Tree for type safety
type GenericBST[T comparable] struct {
	root *GenericTreeNode[T]
	size int
	less func(a, b T) bool
}

type GenericTreeNode[T comparable] struct {
	Data  T
	Left  *GenericTreeNode[T]
	Right *GenericTreeNode[T]
}

// NewGenericBST creates a new generic BST with custom comparison function
func NewGenericBST[T comparable](less func(a, b T) bool) *GenericBST[T] {
	return &GenericBST[T]{
		root: nil,
		size: 0,
		less: less,
	}
}

// Insert adds a value to the generic BST
func (gbst *GenericBST[T]) Insert(data T) {
	gbst.root = gbst.insertNode(gbst.root, data)
	gbst.size++
}

func (gbst *GenericBST[T]) insertNode(node *GenericTreeNode[T], data T) *GenericTreeNode[T] {
	if node == nil {
		return &GenericTreeNode[T]{Data: data}
	}

	if gbst.less(data, node.Data) {
		node.Left = gbst.insertNode(node.Left, data)
	} else if gbst.less(node.Data, data) {
		node.Right = gbst.insertNode(node.Right, data)
	}
	// Equal values - no insertion for duplicates

	return node
}

// Search finds a value in the generic BST
func (gbst *GenericBST[T]) Search(data T) bool {
	return gbst.searchNode(gbst.root, data)
}

func (gbst *GenericBST[T]) searchNode(node *GenericTreeNode[T], data T) bool {
	if node == nil {
		return false
	}

	if data == node.Data {
		return true
	} else if gbst.less(data, node.Data) {
		return gbst.searchNode(node.Left, data)
	}

	return gbst.searchNode(node.Right, data)
}

// Size returns the number of nodes
func (gbst *GenericBST[T]) Size() int {
	return gbst.size
}

// IsEmpty checks if the tree is empty
func (gbst *GenericBST[T]) IsEmpty() bool {
	return gbst.root == nil
}
