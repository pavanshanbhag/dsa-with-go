package medium

import (
	"fmt"
	"math"
	"strings"
)

// ====================================================================
// PROBLEM 1: Add Two Numbers
// You are given two non-empty linked lists representing two non-negative
// integers stored in reverse order.
// ====================================================================

type ListNode struct {
	Val  int
	Next *ListNode
}

// AddTwoNumbers adds two numbers represented as linked lists
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy
	carry := 0

	for l1 != nil || l2 != nil || carry > 0 {
		sum := carry

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		current.Next = &ListNode{Val: sum % 10}
		carry = sum / 10
		current = current.Next
	}

	return dummy.Next
}

// ====================================================================
// PROBLEM 2: Remove Nth Node From End
// Given the head of a linked list, remove the nth node from end.
// ====================================================================

// RemoveNthFromEnd removes nth node from end using two pointers
func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	slow := dummy
	fast := dummy

	// Move fast pointer n+1 steps ahead
	for i := 0; i <= n; i++ {
		fast = fast.Next
	}

	// Move both pointers until fast reaches end
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}

	// Remove the nth node
	slow.Next = slow.Next.Next

	return dummy.Next
}

// ====================================================================
// PROBLEM 3: Binary Tree Level Order Traversal
// Given the root of a binary tree, return level order traversal.
// ====================================================================

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// LevelOrder performs level order traversal using BFS
func LevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	result := [][]int{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		currentLevel := []int{}

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			currentLevel = append(currentLevel, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		result = append(result, currentLevel)
	}

	return result
}

// ====================================================================
// PROBLEM 4: Validate Binary Search Tree
// Given the root of a binary tree, determine if it is a valid BST.
// ====================================================================

// IsValidBST validates BST using bounds checking
func IsValidBST(root *TreeNode) bool {
	return validateBST(root, math.MinInt64, math.MaxInt64)
}

func validateBST(node *TreeNode, minVal, maxVal int) bool {
	if node == nil {
		return true
	}

	if node.Val <= minVal || node.Val >= maxVal {
		return false
	}

	return validateBST(node.Left, minVal, node.Val) &&
		validateBST(node.Right, node.Val, maxVal)
}

// IsValidBSTInorder validates BST using inorder traversal
func IsValidBSTInorder(root *TreeNode) bool {
	values := []int{}
	inorderTraversal(root, &values)

	for i := 1; i < len(values); i++ {
		if values[i] <= values[i-1] {
			return false
		}
	}

	return true
}

func inorderTraversal(node *TreeNode, values *[]int) {
	if node == nil {
		return
	}

	inorderTraversal(node.Left, values)
	*values = append(*values, node.Val)
	inorderTraversal(node.Right, values)
}

// ====================================================================
// PROBLEM 5: Construct Binary Tree from Preorder and Inorder
// Given preorder and inorder traversal, construct the binary tree.
// ====================================================================

// BuildTree constructs binary tree from preorder and inorder traversals
func BuildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	// Create map for quick inorder index lookup
	inorderMap := make(map[int]int)
	for i, val := range inorder {
		inorderMap[val] = i
	}

	return buildTreeHelper(preorder, 0, len(preorder)-1, inorder, 0, len(inorder)-1, inorderMap)
}

func buildTreeHelper(preorder []int, preStart, preEnd int, inorder []int, inStart, inEnd int, inorderMap map[int]int) *TreeNode {
	if preStart > preEnd || inStart > inEnd {
		return nil
	}

	// Root is first element in preorder
	rootVal := preorder[preStart]
	root := &TreeNode{Val: rootVal}

	// Find root position in inorder
	rootIndex := inorderMap[rootVal]
	leftSubtreeSize := rootIndex - inStart

	// Recursively build left and right subtrees
	root.Left = buildTreeHelper(preorder, preStart+1, preStart+leftSubtreeSize,
		inorder, inStart, rootIndex-1, inorderMap)
	root.Right = buildTreeHelper(preorder, preStart+leftSubtreeSize+1, preEnd,
		inorder, rootIndex+1, inEnd, inorderMap)

	return root
}

// ====================================================================
// PROBLEM 6: Kth Smallest Element in BST
// Given the root of a BST and integer k, return kth smallest value.
// ====================================================================

// KthSmallest finds kth smallest element using inorder traversal
func KthSmallest(root *TreeNode, k int) int {
	count := 0
	result := 0

	var inorder func(*TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil || count >= k {
			return
		}

		inorder(node.Left)

		count++
		if count == k {
			result = node.Val
			return
		}

		inorder(node.Right)
	}

	inorder(root)
	return result
}

// KthSmallestIterative using iterative inorder traversal
func KthSmallestIterative(root *TreeNode, k int) int {
	stack := []*TreeNode{}
	current := root
	count := 0

	for current != nil || len(stack) > 0 {
		// Go to leftmost node
		for current != nil {
			stack = append(stack, current)
			current = current.Left
		}

		// Pop and process
		current = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		count++
		if count == k {
			return current.Val
		}

		current = current.Right
	}

	return -1 // Should not reach here for valid input
}

// ====================================================================
// PROBLEM 7: Lowest Common Ancestor of Binary Tree
// Given a binary tree, find the LCA of two given nodes.
// ====================================================================

// LowestCommonAncestor finds LCA using recursive approach
func LowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}

	left := LowestCommonAncestor(root.Left, p, q)
	right := LowestCommonAncestor(root.Right, p, q)

	if left != nil && right != nil {
		return root // Current node is LCA
	}

	if left != nil {
		return left
	}
	return right
}

// ====================================================================
// PROBLEM 8: Binary Tree Zigzag Level Order Traversal
// Return zigzag level order traversal of binary tree.
// ====================================================================

// ZigzagLevelOrder performs zigzag level order traversal
func ZigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	result := [][]int{}
	queue := []*TreeNode{root}
	leftToRight := true

	for len(queue) > 0 {
		levelSize := len(queue)
		currentLevel := make([]int, levelSize)

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			// Determine position based on direction
			var pos int
			if leftToRight {
				pos = i
			} else {
				pos = levelSize - 1 - i
			}
			currentLevel[pos] = node.Val

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		result = append(result, currentLevel)
		leftToRight = !leftToRight
	}

	return result
}

// ====================================================================
// PROBLEM 9: Design HashMap
// Design a HashMap without using built-in hash table libraries.
// ====================================================================

const BUCKET_SIZE = 1000

type MyHashMap struct {
	buckets [][]Pair
}

type Pair struct {
	key   int
	value int
}

func ConstructorHashMap() MyHashMap {
	return MyHashMap{
		buckets: make([][]Pair, BUCKET_SIZE),
	}
}

func (h *MyHashMap) hash(key int) int {
	return key % BUCKET_SIZE
}

func (h *MyHashMap) Put(key int, value int) {
	bucketIndex := h.hash(key)

	// Check if key already exists
	for i := range h.buckets[bucketIndex] {
		if h.buckets[bucketIndex][i].key == key {
			h.buckets[bucketIndex][i].value = value
			return
		}
	}

	// Add new key-value pair
	h.buckets[bucketIndex] = append(h.buckets[bucketIndex], Pair{key, value})
}

func (h *MyHashMap) Get(key int) int {
	bucketIndex := h.hash(key)

	for _, pair := range h.buckets[bucketIndex] {
		if pair.key == key {
			return pair.value
		}
	}

	return -1 // Key not found
}

func (h *MyHashMap) Remove(key int) {
	bucketIndex := h.hash(key)

	for i, pair := range h.buckets[bucketIndex] {
		if pair.key == key {
			// Remove element by swapping with last and reducing slice
			h.buckets[bucketIndex][i] = h.buckets[bucketIndex][len(h.buckets[bucketIndex])-1]
			h.buckets[bucketIndex] = h.buckets[bucketIndex][:len(h.buckets[bucketIndex])-1]
			return
		}
	}
}

// ====================================================================
// PROBLEM 10: Implement Trie (Prefix Tree)
// Implement a trie with insert, search, and startsWith methods.
// ====================================================================

type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

type Trie struct {
	root *TrieNode
}

func ConstructorTrie() Trie {
	return Trie{
		root: &TrieNode{
			children: make(map[rune]*TrieNode),
			isEnd:    false,
		},
	}
}

func (t *Trie) Insert(word string) {
	current := t.root

	for _, char := range word {
		if _, exists := current.children[char]; !exists {
			current.children[char] = &TrieNode{
				children: make(map[rune]*TrieNode),
				isEnd:    false,
			}
		}
		current = current.children[char]
	}

	current.isEnd = true
}

func (t *Trie) Search(word string) bool {
	current := t.root

	for _, char := range word {
		if _, exists := current.children[char]; !exists {
			return false
		}
		current = current.children[char]
	}

	return current.isEnd
}

func (t *Trie) StartsWith(prefix string) bool {
	current := t.root

	for _, char := range prefix {
		if _, exists := current.children[char]; !exists {
			return false
		}
		current = current.children[char]
	}

	return true
}

// ====================================================================
// HELPER FUNCTIONS
// ====================================================================

// CreateLinkedList creates a linked list from slice
func CreateLinkedList(values []int) *ListNode {
	if len(values) == 0 {
		return nil
	}

	head := &ListNode{Val: values[0]}
	current := head

	for i := 1; i < len(values); i++ {
		current.Next = &ListNode{Val: values[i]}
		current = current.Next
	}

	return head
}

// LinkedListToSlice converts linked list to slice
func LinkedListToSlice(head *ListNode) []int {
	result := []int{}
	current := head

	for current != nil {
		result = append(result, current.Val)
		current = current.Next
	}

	return result
}

// CreateBinaryTree creates binary tree from level-order array
func CreateBinaryTree(values []*int) *TreeNode {
	if len(values) == 0 || values[0] == nil {
		return nil
	}

	root := &TreeNode{Val: *values[0]}
	queue := []*TreeNode{root}
	i := 1

	for len(queue) > 0 && i < len(values) {
		node := queue[0]
		queue = queue[1:]

		// Left child
		if i < len(values) && values[i] != nil {
			node.Left = &TreeNode{Val: *values[i]}
			queue = append(queue, node.Left)
		}
		i++

		// Right child
		if i < len(values) && values[i] != nil {
			node.Right = &TreeNode{Val: *values[i]}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}

// IntPtr creates int pointer
func IntPtr(val int) *int {
	return &val
}

// ====================================================================
// PROBLEM DEMONSTRATION
// ====================================================================

// DemonstrateMediumDataStructures shows all medium data structure problems
func DemonstrateMediumDataStructures() {
	fmt.Println("🎯 Medium Data Structure Problems")
	fmt.Println("=================================")

	// Problem 1: Add Two Numbers
	fmt.Println("\n1. Add Two Numbers:")
	l1 := CreateLinkedList([]int{2, 4, 3}) // represents 342
	l2 := CreateLinkedList([]int{5, 6, 4}) // represents 465
	sum := AddTwoNumbers(l1, l2)
	fmt.Printf("342 + 465 = %v (reversed)\n", LinkedListToSlice(sum)) // should be [7,0,8] for 807

	// Problem 2: Remove Nth From End
	fmt.Println("\n2. Remove Nth Node From End:")
	list := CreateLinkedList([]int{1, 2, 3, 4, 5})
	fmt.Printf("Original: %v\n", LinkedListToSlice(list))
	list = CreateLinkedList([]int{1, 2, 3, 4, 5}) // recreate since we modified it
	removed := RemoveNthFromEnd(list, 2)
	fmt.Printf("Remove 2nd from end: %v\n", LinkedListToSlice(removed))

	// Problem 3: Level Order Traversal
	fmt.Println("\n3. Binary Tree Level Order Traversal:")
	tree := CreateBinaryTree([]*int{IntPtr(3), IntPtr(9), IntPtr(20), nil, nil, IntPtr(15), IntPtr(7)})
	levels := LevelOrder(tree)
	fmt.Printf("Level order: %v\n", levels)

	// Problem 4: Validate BST
	fmt.Println("\n4. Validate Binary Search Tree:")
	bst := CreateBinaryTree([]*int{IntPtr(2), IntPtr(1), IntPtr(3)})
	invalidBst := CreateBinaryTree([]*int{IntPtr(5), IntPtr(1), IntPtr(4), nil, nil, IntPtr(3), IntPtr(6)})
	fmt.Printf("Tree [2,1,3] is valid BST: %t\n", IsValidBST(bst))
	fmt.Printf("Tree [5,1,4,null,null,3,6] is valid BST: %t\n", IsValidBST(invalidBst))

	// Problem 5: Build Tree
	fmt.Println("\n5. Construct Binary Tree:")
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	constructed := BuildTree(preorder, inorder)
	constructedLevels := LevelOrder(constructed)
	fmt.Printf("Built from preorder %v and inorder %v:\n", preorder, inorder)
	fmt.Printf("Level order: %v\n", constructedLevels)

	// Problem 6: Kth Smallest in BST
	fmt.Println("\n6. Kth Smallest Element in BST:")
	bst2 := CreateBinaryTree([]*int{IntPtr(3), IntPtr(1), IntPtr(4), nil, IntPtr(2)})
	k := 1
	kthSmallest := KthSmallest(bst2, k)
	fmt.Printf("%dst smallest element: %d\n", k, kthSmallest)

	// Problem 7: Lowest Common Ancestor
	fmt.Println("\n7. Lowest Common Ancestor:")
	_ = CreateBinaryTree([]*int{IntPtr(3), IntPtr(5), IntPtr(1), IntPtr(6), IntPtr(2), IntPtr(0), IntPtr(8), nil, nil, IntPtr(7), IntPtr(4)})
	// For demonstration, we'll assume nodes with values 5 and 1
	fmt.Printf("Tree structure created for LCA demonstration\n")

	// Problem 8: Zigzag Level Order
	fmt.Println("\n8. Zigzag Level Order Traversal:")
	zigzagTree := CreateBinaryTree([]*int{IntPtr(3), IntPtr(9), IntPtr(20), nil, nil, IntPtr(15), IntPtr(7)})
	zigzag := ZigzagLevelOrder(zigzagTree)
	fmt.Printf("Zigzag traversal: %v\n", zigzag)

	// Problem 9: Design HashMap
	fmt.Println("\n9. Design HashMap:")
	hashMap := ConstructorHashMap()
	hashMap.Put(1, 1)
	hashMap.Put(2, 2)
	fmt.Printf("Get(1): %d\n", hashMap.Get(1))
	fmt.Printf("Get(3): %d\n", hashMap.Get(3)) // should return -1
	hashMap.Put(2, 1)                          // update existing key
	fmt.Printf("Get(2) after update: %d\n", hashMap.Get(2))
	hashMap.Remove(2)
	fmt.Printf("Get(2) after removal: %d\n", hashMap.Get(2))

	// Problem 10: Implement Trie
	fmt.Println("\n10. Implement Trie:")
	trie := ConstructorTrie()
	trie.Insert("apple")
	fmt.Printf("Search 'apple': %t\n", trie.Search("apple"))
	fmt.Printf("Search 'app': %t\n", trie.Search("app"))
	fmt.Printf("StartsWith 'app': %t\n", trie.StartsWith("app"))
	trie.Insert("app")
	fmt.Printf("Search 'app' after insert: %t\n", trie.Search("app"))
}

// ProblemComplexityAnalysis provides complexity analysis for medium problems
func ProblemComplexityAnalysis() {
	fmt.Println("\n📊 Medium Data Structures Complexity Analysis")
	fmt.Println("=============================================")

	problems := []struct {
		name     string
		timeOpt  string
		spaceOpt string
		approach string
	}{
		{"Add Two Numbers", "O(max(m,n))", "O(max(m,n))", "Linked List Traversal"},
		{"Remove Nth From End", "O(n)", "O(1)", "Two Pointers"},
		{"Level Order Traversal", "O(n)", "O(w)", "BFS Queue"},
		{"Validate BST", "O(n)", "O(h)", "Recursive Bounds"},
		{"Build Tree", "O(n)", "O(n)", "Divide & Conquer"},
		{"Kth Smallest in BST", "O(h+k)", "O(h)", "Inorder Traversal"},
		{"Lowest Common Ancestor", "O(n)", "O(h)", "Recursive Search"},
		{"Zigzag Level Order", "O(n)", "O(w)", "BFS with Direction"},
		{"Design HashMap", "O(1) average", "O(n)", "Hash with Chaining"},
		{"Implement Trie", "O(m)", "O(ALPHABET_SIZE×N×M)", "Tree Structure"},
	}

	fmt.Printf("%-25s %-15s %-15s %-25s\n", "Problem", "Time", "Space", "Best Approach")
	fmt.Println(strings.Repeat("-", 85))

	for _, p := range problems {
		fmt.Printf("%-25s %-15s %-15s %-25s\n", p.name, p.timeOpt, p.spaceOpt, p.approach)
	}
}
