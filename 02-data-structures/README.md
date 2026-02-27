# Data Structures in Go

This module covers the implementation and analysis of data structures in Go, organized by their fundamental properties and characteristics.

## 📁 Module Structure

```
02-data-structures/
├── README.md                 # This file
├── linear/                   # Sequential data structures
│   ├── README.md
│   ├── arrays.go             # DynamicArray, GenericDynamicArray
│   ├── stack.go              # ArrayStack, LinkedStack, SafeStack, GenericStack
│   ├── queue.go              # ArrayQueue, DynamicArrayQueue, LinkedQueue, PriorityQueue, ChannelQueue, GenericQueue, SafeQueue
│   ├── linked_list.go        # Singly, Doubly, Circular, GenericLinkedList
│   ├── linear_test.go        # Tests for arrays and stacks
│   ├── linked_list_test.go   # Tests for linked lists
│   └── queue_linkedlist_test.go  # Tests for queues and linked list apps
└── non-linear/               # Hierarchical data structures
    ├── README.md
    ├── tree.go               # BinaryTree, BinarySearchTree, MinHeap, GenericBST
    └── tree_test.go          # Tests for BST, heap, binary tree
```

## 🧪 How to Verify

From the **repository root**:

```bash
# Run all data structure tests
go test ./02-data-structures/linear/ ./02-data-structures/non-linear/ -v

# Run with benchmarks
go test ./02-data-structures/linear/ ./02-data-structures/non-linear/ -bench=. -benchmem
```

## 📚 Organization

### Linear Data Structures (`/linear`)
Data structures where elements are arranged in a sequential manner, each element connected to its previous and next element.

- **Arrays & Dynamic Arrays**: Fixed and resizable sequential containers
- **Linked Lists**: Node-based sequential structures (singly, doubly, circular)
- **Stacks**: LIFO (Last In, First Out) data structure
- **Queues**: FIFO (First In, First Out) data structure with variants

### Non-Linear Data Structures (`/non-linear`)
Data structures where elements are arranged in a hierarchical manner or complex relationships.

- **Binary Trees**: Hierarchical tree structures with up to 2 children per node
- **Binary Search Trees**: Ordered binary trees for efficient searching
- **Heaps**: Complete binary trees with heap property (min/max)

## 🎯 Learning Objectives

- Understand when to use each data structure
- Implement thread-safe versions using Go's concurrency primitives
- Analyze time and space complexity for each operation
- Apply data structures to solve real-world problems
- Master Go-specific optimizations and idioms

## 🧪 Testing & Validation

Each data structure includes:
- **Complete Implementation**: Built from scratch with proper Go idioms
- **Comprehensive Test Suite**: Unit tests covering all operations and edge cases
- **Benchmark Comparisons**: Performance analysis between different implementations
- **Real-world Examples**: Practical applications and use cases
- **Error Handling**: Proper error handling following Go conventions

## 🚀 Go-Specific Features

- **Interface-based Design**: Clean abstractions and polymorphism
- **Generic Implementations**: Type-safe collections using Go 1.24+ generics
- **Memory Efficiency**: Zero-allocation techniques where possible
- **Concurrent Safety**: Mutex-protected versions for multi-threaded access
- **Channel Integration**: Go-idiomatic implementations using channels

## 📈 Performance Insights

### Linear Structures Comparison:
- **Arrays**: Best cache locality, O(1) access by index
- **Linked Lists**: Dynamic sizing, O(1) insertion/deletion at known positions
- **Stacks**: Optimized for LIFO operations, function call management
- **Queues**: Optimized for FIFO operations, task scheduling, BFS algorithms

### Non-Linear Structures Comparison:
- **BST**: O(log n) average search/insert/delete, can degrade to O(n)
- **Heaps**: O(log n) insert/extract, O(1) peek, excellent for priority queues
- **Generic Trees**: Flexible structure for hierarchical data

## 🔄 Cross-Package Usage

Some structures depend on others:
- Trees use simple queues for level-order traversal
- Priority queues internally use heaps
- Various algorithms use stacks and queues

## 📝 Next Steps

After mastering these data structures, you'll be ready for:
- **Graph Algorithms**: BFS, DFS, shortest paths, MST
- **Advanced Trees**: AVL, Red-Black, B-Trees
- **Hash Tables**: Collision resolution, load factor management
- **String Algorithms**: Tries, suffix trees, pattern matching
- **Error Handling**: Idiomatic Go error patterns
