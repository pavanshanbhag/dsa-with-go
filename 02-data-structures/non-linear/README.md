# Non-Linear Data Structures

Hierarchical and interconnected data structures where elements have complex relationships beyond simple sequential ordering.

## 🌳 Structures Implemented

### Binary Trees (`tree.go`)
- **BinaryTree**: General binary tree with traversal operations
- **TreeNode**: Basic node structure with left/right children
- **Features**: Inorder, preorder, postorder, level-order traversals
- **Applications**: Expression trees, decision trees, file systems

### Binary Search Trees (`tree.go`)
- **BinarySearchTree**: Ordered binary tree for efficient searching
- **Operations**: Insert, search, delete with BST property maintenance
- **Features**: All tree traversals, size tracking, empty tree handling
- **Time Complexity**: O(log n) average case, O(n) worst case
- **Applications**: Database indexing, sorted data maintenance

### Heaps (`tree.go`)
- **MinHeap**: Complete binary tree with min-heap property
- **Operations**: Insert, extract-min, peek with heap property maintenance
- **Implementation**: Array-based representation for efficiency
- **Time Complexity**: O(log n) insert/extract, O(1) peek
- **Applications**: Priority queues, heap sort, scheduling algorithms

### Generic Trees (`tree.go`)
- **GenericBST**: Type-safe binary search tree using generics
- **Flexible Comparison**: Custom comparison functions for any comparable type
- **Features**: Type safety, custom ordering, performance optimization

## 🧪 Testing

Comprehensive test coverage in `tree_test.go`:
- **TestBinarySearchTree**: Full BST functionality testing
- **TestMinHeap**: Heap property validation and operations
- **TestBinaryTree**: Basic tree operations and traversals

### How to run (from repo root)
```bash
go test ./02-data-structures/non-linear/ -v
go test ./02-data-structures/non-linear/ -bench=. -benchmem
```

### Test Coverage:
- ✅ Basic operations (insert, delete, search, extract)
- ✅ Tree traversals (all four types)
- ✅ Heap property maintenance
- ✅ Edge cases (empty trees, single nodes)
- ✅ Error conditions and boundary testing
- ✅ Performance benchmarks

## 🚀 Performance Characteristics

| Structure | Insert | Delete | Search | Find Min/Max | Space |
|-----------|---------|---------|---------|---------|---------|
| BST | O(log n)* | O(log n)* | O(log n)* | O(log n)* | O(n) |
| Min Heap | O(log n) | O(log n) | O(n) | O(1) | O(n) |
| Binary Tree | O(n) | O(n) | O(n) | O(n) | O(n) |

*Average case - worst case can be O(n) for unbalanced trees

## 💡 When to Use Each

### Binary Search Trees
- Need sorted order maintenance
- Frequent search operations
- Range queries required
- Dynamic data with insertions/deletions

### Heaps (Priority Queues)
- Task scheduling by priority
- Finding minimum/maximum efficiently
- Heap sort algorithm
- Merging k sorted arrays

### General Binary Trees
- Hierarchical data representation
- Expression parsing
- Decision tree algorithms
- File system structures

## 🔄 Tree Traversals

All tree structures support four traversal methods:

### Inorder (Left → Root → Right)
```go
result := tree.InorderTraversal()
// For BST: returns sorted order
```

### Preorder (Root → Left → Right)
```go
result := tree.PreorderTraversal()
// Useful for tree serialization
```

### Postorder (Left → Right → Root)
```go
result := tree.PostorderTraversal()
// Useful for tree deletion
```

### Level Order (Breadth-First)
```go
result := tree.LevelOrderTraversal()
// Prints tree level by level
```

## 🎯 Key Features

### Heap Implementation Details
- **Array-based**: Efficient memory usage and cache performance
- **Parent-child relationship**: `parent = (i-1)/2`, `left = 2*i+1`, `right = 2*i+2`
- **Heapify operations**: Maintain heap property after modifications
- **Complete tree**: All levels filled except possibly the last

### BST Implementation Details
- **Recursive operations**: Clean and intuitive implementation
- **In-order successor**: For deletion of nodes with two children
- **Size tracking**: O(1) size queries
- **Flexible comparison**: Support for custom types

## 🔧 Go-Specific Optimizations

- **Interface{} vs Generics**: Both approaches provided for compatibility
- **Error handling**: Proper Go error patterns for edge cases
- **Memory efficiency**: Array-based heap implementation
- **Slice-based queue**: Simple internal queue for level-order traversal
- **Method receivers**: Consistent pointer receivers for efficiency

## 📈 Performance Tips

1. **For BST**: Keep tree balanced for optimal performance
2. **For Heaps**: Array-based implementation is more cache-friendly than pointer-based
3. **For Traversals**: Choose the right traversal for your use case
4. **Memory**: Consider tree depth for stack-based recursive operations

## 🚧 Future Enhancements

These implementations provide a solid foundation. Consider adding:
- **Self-balancing trees**: AVL, Red-Black trees
- **B-trees**: For database applications
- **Trie**: For string algorithms
- **Segment trees**: For range queries
