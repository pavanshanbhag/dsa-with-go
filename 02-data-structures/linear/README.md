# Linear Data Structures

Sequential data structures where elements have a linear relationship - each element (except first and last) has exactly one predecessor and one successor.

## 📊 Structures Implemented

### Arrays & Dynamic Arrays (`arrays.go`)
- **DynamicArray**: Resizable array with capacity management
- **GenericDynamicArray**: Type-safe version using generics
- **Features**: Automatic resizing, shrinking, iterator pattern
- **Time Complexity**: O(1) access, O(1) amortized append, O(n) insert/delete

### Stacks (`stack.go`)
- **ArrayStack**: Array-based LIFO implementation
- **LinkedStack**: Linked list-based implementation
- **SafeStack**: Thread-safe version with mutex
- **GenericStack**: Type-safe generic implementation
- **Applications**: Parentheses matching, postfix evaluation
- **Time Complexity**: O(1) for all operations

### Queues (`queue.go`)
- **ArrayQueue**: Fixed-size circular buffer
- **DynamicArrayQueue**: Resizable array-based queue
- **LinkedQueue**: Linked list-based queue
- **PriorityQueue**: Heap-based priority queue
- **ChannelQueue**: Go-idiomatic concurrent queue using channels
- **GenericQueue**: Type-safe generic implementation
- **Time Complexity**: O(1) enqueue/dequeue (except priority queue: O(log n))

### Linked Lists (`linked_list.go`)
- **SinglyLinkedList**: Basic linked list with forward traversal
- **DoublyLinkedList**: Bidirectional linked list
- **CircularLinkedList**: Circular linked list for round-robin algorithms
- **Features**: Cycle detection, list reversal, middle element finding
- **Time Complexity**: O(1) insert/delete at known position, O(n) search

## 🧪 Testing

Each structure has comprehensive tests in:
- `linear_test.go`: Tests for arrays and stacks
- `linked_list_test.go`: Tests for all linked list variants
- `queue_linkedlist_test.go`: Additional queue-specific tests

### Test Coverage:
- ✅ Basic operations (insert, delete, search)
- ✅ Edge cases (empty structures, single elements)
- ✅ Error conditions
- ✅ Performance benchmarks
- ✅ Memory usage patterns
- ✅ Thread safety (where applicable)

## 🚀 Performance Characteristics

| Structure | Access | Insert | Delete | Search | Memory |
|-----------|---------|---------|---------|---------|---------|
| Dynamic Array | O(1) | O(1)* | O(n) | O(n) | Compact |
| Stack | O(1) top | O(1) | O(1) | O(n) | Minimal |
| Queue | O(1) ends | O(1) | O(1) | O(n) | Efficient |
| Linked List | O(n) | O(1)** | O(1)** | O(n) | Overhead |

*Amortized, **At known position

## 💡 When to Use Each

### Dynamic Arrays
- Random access needed
- Cache-friendly iterations
- Memory usage is a concern

### Stacks
- Function call management
- Expression evaluation
- Backtracking algorithms
- Undo operations

### Queues
- Task scheduling
- BFS algorithms
- Producer-consumer patterns
- Request buffering

### Linked Lists
- Frequent insertions/deletions
- Unknown or highly variable size
- Memory is fragmented
- Implementing other structures

## 🔧 Go-Specific Optimizations

- **Slices over arrays**: Leveraging Go's slice capabilities
- **Interface-based design**: Clean abstractions
- **Channel queues**: Leveraging Go's concurrency primitives
- **Generics**: Type safety with zero-cost abstractions
- **Error handling**: Proper Go error patterns
- **Memory efficiency**: Minimize allocations where possible
