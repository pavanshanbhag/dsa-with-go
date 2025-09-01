package easy

import (
	"fmt"
	"strings"
)

// ====================================================================
// PROBLEM 1: Valid Parentheses
// Given a string containing just the characters '(', ')', '{', '}', '[' and ']',
// determine if the input string is valid.
// ====================================================================

// IsValid checks if parentheses are properly balanced using stack
func IsValid(s string) bool {
	stack := []rune{}
	pairs := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, char := range s {
		switch char {
		case '(', '{', '[':
			// Opening bracket - push to stack
			stack = append(stack, char)
		case ')', '}', ']':
			// Closing bracket - check matching
			if len(stack) == 0 {
				return false
			}
			if stack[len(stack)-1] != pairs[char] {
				return false
			}
			// Pop from stack
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

// ====================================================================
// PROBLEM 2: Implement Queue using Stacks
// Implement a first in first out (FIFO) queue using only two stacks.
// ====================================================================

type MyQueue struct {
	inStack  []int
	outStack []int
}

func Constructor() MyQueue {
	return MyQueue{
		inStack:  []int{},
		outStack: []int{},
	}
}

func (q *MyQueue) Push(x int) {
	q.inStack = append(q.inStack, x)
}

func (q *MyQueue) Pop() int {
	q.ensureOutStack()
	if len(q.outStack) == 0 {
		return -1 // Invalid operation
	}

	result := q.outStack[len(q.outStack)-1]
	q.outStack = q.outStack[:len(q.outStack)-1]
	return result
}

func (q *MyQueue) Peek() int {
	q.ensureOutStack()
	if len(q.outStack) == 0 {
		return -1 // Invalid operation
	}

	return q.outStack[len(q.outStack)-1]
}

func (q *MyQueue) Empty() bool {
	return len(q.inStack) == 0 && len(q.outStack) == 0
}

func (q *MyQueue) ensureOutStack() {
	if len(q.outStack) == 0 {
		// Transfer all elements from inStack to outStack
		for len(q.inStack) > 0 {
			q.outStack = append(q.outStack, q.inStack[len(q.inStack)-1])
			q.inStack = q.inStack[:len(q.inStack)-1]
		}
	}
}

// ====================================================================
// PROBLEM 3: Binary Tree Inorder Traversal
// Given the root of a binary tree, return the inorder traversal of its nodes' values.
// ====================================================================

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// InorderTraversal using recursion
func InorderTraversal(root *TreeNode) []int {
	result := []int{}
	inorderHelper(root, &result)
	return result
}

func inorderHelper(node *TreeNode, result *[]int) {
	if node == nil {
		return
	}

	inorderHelper(node.Left, result)
	*result = append(*result, node.Val)
	inorderHelper(node.Right, result)
}

// InorderTraversalIterative using stack
func InorderTraversalIterative(root *TreeNode) []int {
	result := []int{}
	stack := []*TreeNode{}
	current := root

	for current != nil || len(stack) > 0 {
		// Go to leftmost node
		for current != nil {
			stack = append(stack, current)
			current = current.Left
		}

		// Pop and process
		current = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, current.Val)

		// Move to right subtree
		current = current.Right
	}

	return result
}

// ====================================================================
// PROBLEM 4: Maximum Depth of Binary Tree
// Given the root of a binary tree, return its maximum depth.
// ====================================================================

// MaxDepth calculates maximum depth using recursion
func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth := MaxDepth(root.Left)
	rightDepth := MaxDepth(root.Right)

	return 1 + max(leftDepth, rightDepth)
}

// MaxDepthIterative using level-order traversal
func MaxDepthIterative(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := []*TreeNode{root}
	depth := 0

	for len(queue) > 0 {
		levelSize := len(queue)
		depth++

		// Process all nodes at current level
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}

	return depth
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// ====================================================================
// PROBLEM 5: Same Tree
// Given the roots of two binary trees, check if they are the same.
// ====================================================================

// IsSameTree checks if two trees are identical
func IsSameTree(p *TreeNode, q *TreeNode) bool {
	// Both null
	if p == nil && q == nil {
		return true
	}

	// One null, one not null
	if p == nil || q == nil {
		return false
	}

	// Both not null - check value and subtrees
	return p.Val == q.Val &&
		IsSameTree(p.Left, q.Left) &&
		IsSameTree(p.Right, q.Right)
}

// ====================================================================
// PROBLEM 6: Symmetric Tree
// Given the root of a binary tree, check whether it is a mirror of itself.
// ====================================================================

// IsSymmetric checks if tree is symmetric
func IsSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isMirror(root.Left, root.Right)
}

func isMirror(left, right *TreeNode) bool {
	// Both null
	if left == nil && right == nil {
		return true
	}

	// One null, one not null
	if left == nil || right == nil {
		return false
	}

	// Both not null - check value and mirror subtrees
	return left.Val == right.Val &&
		isMirror(left.Left, right.Right) &&
		isMirror(left.Right, right.Left)
}

// ====================================================================
// PROBLEM 7: Linked List Cycle
// Given head of a linked list, determine if it has a cycle.
// ====================================================================

type ListNode struct {
	Val  int
	Next *ListNode
}

// HasCycle detects cycle using Floyd's algorithm (tortoise and hare)
func HasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow := head
	fast := head.Next

	for fast != nil && fast.Next != nil {
		if slow == fast {
			return true
		}
		slow = slow.Next
		fast = fast.Next.Next
	}

	return false
}

// HasCycleHashSet detects cycle using hash set
func HasCycleHashSet(head *ListNode) bool {
	visited := make(map[*ListNode]bool)

	current := head
	for current != nil {
		if visited[current] {
			return true
		}
		visited[current] = true
		current = current.Next
	}

	return false
}

// ====================================================================
// PROBLEM 8: Merge Two Sorted Lists
// Merge two sorted linked lists and return as a new sorted list.
// ====================================================================

// MergeTwoLists merges two sorted linked lists
func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}
		current = current.Next
	}

	// Append remaining nodes
	if list1 != nil {
		current.Next = list1
	} else {
		current.Next = list2
	}

	return dummy.Next
}

// MergeTwoListsRecursive using recursion
func MergeTwoListsRecursive(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	if list1.Val <= list2.Val {
		list1.Next = MergeTwoListsRecursive(list1.Next, list2)
		return list1
	} else {
		list2.Next = MergeTwoListsRecursive(list1, list2.Next)
		return list2
	}
}

// ====================================================================
// PROBLEM 9: Remove Duplicates from Sorted List
// Given the head of a sorted linked list, delete all duplicates.
// ====================================================================

// DeleteDuplicates removes duplicates from sorted linked list
func DeleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	current := head
	for current != nil && current.Next != nil {
		if current.Val == current.Next.Val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}

	return head
}

// ====================================================================
// PROBLEM 10: Reverse Linked List
// Given the head of a singly linked list, reverse the list.
// ====================================================================

// ReverseList reverses linked list iteratively
func ReverseList(head *ListNode) *ListNode {
	var prev *ListNode
	current := head

	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}

	return prev
}

// ReverseListRecursive reverses linked list recursively
func ReverseListRecursive(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// Reverse the rest of the list
	newHead := ReverseListRecursive(head.Next)

	// Reverse current connection
	head.Next.Next = head
	head.Next = nil

	return newHead
}

// ====================================================================
// HELPER FUNCTIONS FOR DEMONSTRATION
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

// LinkedListToSlice converts linked list to slice for display
func LinkedListToSlice(head *ListNode) []int {
	result := []int{}
	current := head

	for current != nil {
		result = append(result, current.Val)
		current = current.Next
	}

	return result
}

// CreateBinaryTree creates binary tree from level-order array (nil for missing nodes)
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

// Helper function to create int pointer
func IntPtr(val int) *int {
	return &val
}

// ====================================================================
// PROBLEM DEMONSTRATION
// ====================================================================

// DemonstrateEasyDataStructures shows all easy data structure problems
func DemonstrateEasyDataStructures() {
	fmt.Println("🎯 Easy Data Structure Problems")
	fmt.Println("===============================")

	// Problem 1: Valid Parentheses
	fmt.Println("\n1. Valid Parentheses:")
	testCases := []string{"()", "()[]{}", "(]", "([)]", "{[]}"}
	for _, test := range testCases {
		result := IsValid(test)
		fmt.Printf("'%s': %t\n", test, result)
	}

	// Problem 2: Queue using Stacks
	fmt.Println("\n2. Queue using Stacks:")
	queue := Constructor()
	queue.Push(1)
	queue.Push(2)
	fmt.Printf("Peek: %d\n", queue.Peek())
	fmt.Printf("Pop: %d\n", queue.Pop())
	fmt.Printf("Empty: %t\n", queue.Empty())

	// Problem 3: Binary Tree Inorder Traversal
	fmt.Println("\n3. Binary Tree Inorder Traversal:")
	// Create tree: [1,null,2,3]
	tree1 := CreateBinaryTree([]*int{IntPtr(1), nil, IntPtr(2), IntPtr(3)})
	inorder := InorderTraversal(tree1)
	inorderIter := InorderTraversalIterative(tree1)
	fmt.Printf("Recursive: %v\n", inorder)
	fmt.Printf("Iterative: %v\n", inorderIter)

	// Problem 4: Maximum Depth
	fmt.Println("\n4. Maximum Depth of Binary Tree:")
	// Create tree: [3,9,20,null,null,15,7]
	tree2 := CreateBinaryTree([]*int{IntPtr(3), IntPtr(9), IntPtr(20), nil, nil, IntPtr(15), IntPtr(7)})
	depth := MaxDepth(tree2)
	depthIter := MaxDepthIterative(tree2)
	fmt.Printf("Recursive depth: %d\n", depth)
	fmt.Printf("Iterative depth: %d\n", depthIter)

	// Problem 5: Same Tree
	fmt.Println("\n5. Same Tree:")
	tree3 := CreateBinaryTree([]*int{IntPtr(1), IntPtr(2), IntPtr(3)})
	tree4 := CreateBinaryTree([]*int{IntPtr(1), IntPtr(2), IntPtr(3)})
	tree5 := CreateBinaryTree([]*int{IntPtr(1), IntPtr(2), IntPtr(1)})
	fmt.Printf("Tree [1,2,3] == Tree [1,2,3]: %t\n", IsSameTree(tree3, tree4))
	fmt.Printf("Tree [1,2,3] == Tree [1,2,1]: %t\n", IsSameTree(tree3, tree5))

	// Problem 6: Symmetric Tree
	fmt.Println("\n6. Symmetric Tree:")
	symTree := CreateBinaryTree([]*int{IntPtr(1), IntPtr(2), IntPtr(2), IntPtr(3), IntPtr(4), IntPtr(4), IntPtr(3)})
	asymTree := CreateBinaryTree([]*int{IntPtr(1), IntPtr(2), IntPtr(2), nil, IntPtr(3), nil, IntPtr(3)})
	fmt.Printf("Symmetric tree: %t\n", IsSymmetric(symTree))
	fmt.Printf("Asymmetric tree: %t\n", IsSymmetric(asymTree))

	// Problem 7: Linked List Cycle
	fmt.Println("\n7. Linked List Cycle:")
	list1 := CreateLinkedList([]int{3, 2, 0, -4})
	list2 := CreateLinkedList([]int{1, 2})
	fmt.Printf("List [3,2,0,-4] has cycle: %t\n", HasCycle(list1))
	fmt.Printf("List [1,2] has cycle: %t\n", HasCycle(list2))

	// Problem 8: Merge Two Sorted Lists
	fmt.Println("\n8. Merge Two Sorted Lists:")
	list3 := CreateLinkedList([]int{1, 2, 4})
	list4 := CreateLinkedList([]int{1, 3, 4})
	merged := MergeTwoLists(list3, list4)
	fmt.Printf("Merge [1,2,4] + [1,3,4]: %v\n", LinkedListToSlice(merged))

	// Problem 9: Remove Duplicates
	fmt.Println("\n9. Remove Duplicates from Sorted List:")
	list5 := CreateLinkedList([]int{1, 1, 2})
	deduplicated := DeleteDuplicates(list5)
	fmt.Printf("Remove duplicates from [1,1,2]: %v\n", LinkedListToSlice(deduplicated))

	// Problem 10: Reverse Linked List
	fmt.Println("\n10. Reverse Linked List:")
	list6 := CreateLinkedList([]int{1, 2, 3, 4, 5})
	original := LinkedListToSlice(list6)
	reversed := ReverseList(list6)
	fmt.Printf("Original: %v\n", original)
	fmt.Printf("Reversed: %v\n", LinkedListToSlice(reversed))
}

// ProblemComplexityAnalysis provides complexity analysis for all problems
func ProblemComplexityAnalysis() {
	fmt.Println("\n📊 Easy Data Structures Complexity Analysis")
	fmt.Println("===========================================")

	problems := []struct {
		name     string
		timeOpt  string
		spaceOpt string
		approach string
	}{
		{"Valid Parentheses", "O(n)", "O(n)", "Stack"},
		{"Queue using Stacks", "O(1) amortized", "O(n)", "Two Stacks"},
		{"Tree Inorder Traversal", "O(n)", "O(h)", "Recursion/Stack"},
		{"Maximum Depth", "O(n)", "O(h)", "DFS Recursion"},
		{"Same Tree", "O(min(m,n))", "O(min(m,n))", "DFS Comparison"},
		{"Symmetric Tree", "O(n)", "O(h)", "Mirror Recursion"},
		{"Linked List Cycle", "O(n)", "O(1)", "Floyd's Algorithm"},
		{"Merge Sorted Lists", "O(m+n)", "O(1)", "Two Pointers"},
		{"Remove Duplicates", "O(n)", "O(1)", "Single Pass"},
		{"Reverse Linked List", "O(n)", "O(1)", "Iterative Reversal"},
	}

	fmt.Printf("%-25s %-15s %-10s %-20s\n", "Problem", "Time", "Space", "Best Approach")
	fmt.Println(strings.Repeat("-", 75))

	for _, p := range problems {
		fmt.Printf("%-25s %-15s %-10s %-20s\n", p.name, p.timeOpt, p.spaceOpt, p.approach)
	}
}
