# Go DSA Quick Reference

## 🚀 Common Patterns & Idioms

### 1. Generic Data Structures (Go 1.24+)

```go
// Type-safe stack
type Stack[T any] struct {
    items []T
}

func (s *Stack[T]) Push(item T) {
    s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, bool) {
    if len(s.items) == 0 {
        var zero T
        return zero, false
    }
    index := len(s.items) - 1
    item := s.items[index]
    s.items = s.items[:index]
    return item, true
}
```

### 2. Interface-Based Design

```go
type Container interface {
    Size() int
    IsEmpty() bool
    Clear()
}

type Stack interface {
    Container
    Push(value interface{})
    Pop() (interface{}, error)
    Peek() (interface{}, error)
}
```

### 3. Thread-Safe Operations

```go
type SafeStack struct {
    items []interface{}
    mutex sync.RWMutex
}

func (s *SafeStack) Push(item interface{}) {
    s.mutex.Lock()
    defer s.mutex.Unlock()
    s.items = append(s.items, item)
}

func (s *SafeStack) Pop() (interface{}, error) {
    s.mutex.Lock()
    defer s.mutex.Unlock()
    // ... pop logic
}
```

### 4. Memory-Efficient Slice Growth

```go
// Pre-allocate when size is known
func NewFixedCapacityArray(capacity int) []int {
    return make([]int, 0, capacity)
}

// Controlled growth to avoid frequent reallocations
func (da *DynamicArray) ensureCapacity(minCapacity int) {
    if minCapacity > da.capacity {
        newCapacity := da.capacity
        for newCapacity < minCapacity {
            if newCapacity < 1024 {
                newCapacity *= 2
            } else {
                newCapacity += newCapacity / 4 // Grow by 25%
            }
        }
        da.resize(newCapacity)
    }
}
```

### 5. Iterator Pattern

```go
type Iterator[T any] interface {
    HasNext() bool
    Next() (T, error)
    Reset()
}

type SliceIterator[T any] struct {
    data  []T
    index int
}

func (it *SliceIterator[T]) HasNext() bool {
    return it.index < len(it.data)
}

func (it *SliceIterator[T]) Next() (T, error) {
    if !it.HasNext() {
        var zero T
        return zero, errors.New("no more elements")
    }
    value := it.data[it.index]
    it.index++
    return value, nil
}
```

## 📊 Complexity Cheat Sheet

### Array/Slice Operations

| Operation | Time | Space | Notes |
|-----------|------|-------|-------|
| Access | O(1) | O(1) | Direct indexing |
| Search | O(n) | O(1) | Linear scan |
| Insert (end) | O(1)* | O(1) | *Amortized |
| Insert (middle) | O(n) | O(1) | Requires shifting |
| Delete (end) | O(1) | O(1) | No shifting needed |
| Delete (middle) | O(n) | O(1) | Requires shifting |

### Stack Operations

| Operation | Time | Space | Notes |
|-----------|------|-------|-------|
| Push | O(1) | O(1) | Top insertion |
| Pop | O(1) | O(1) | Top removal |
| Peek/Top | O(1) | O(1) | Access top element |
| Search | O(n) | O(1) | Not typical stack operation |

### Queue Operations

| Operation | Time | Space | Notes |
|-----------|------|-------|-------|
| Enqueue | O(1) | O(1) | Rear insertion |
| Dequeue | O(1) | O(1) | Front removal |
| Front | O(1) | O(1) | Access front element |
| Search | O(n) | O(1) | Not typical queue operation |

## 🛠 Go-Specific Optimizations

### 1. Avoid Interface{} When Possible

```go
// ❌ Slow - boxing/unboxing overhead
type SlowStack struct {
    items []interface{}
}

// ✅ Fast - no boxing overhead
type FastStack[T any] struct {
    items []T
}
```

### 2. Use Channels for Concurrent Queues

```go
type ChannelQueue[T any] struct {
    ch chan T
}

func NewChannelQueue[T any](capacity int) *ChannelQueue[T] {
    return &ChannelQueue[T]{
        ch: make(chan T, capacity),
    }
}

func (cq *ChannelQueue[T]) Enqueue(item T) error {
    select {
    case cq.ch <- item:
        return nil
    default:
        return errors.New("queue is full")
    }
}

func (cq *ChannelQueue[T]) Dequeue() (T, error) {
    select {
    case item := <-cq.ch:
        return item, nil
    default:
        var zero T
        return zero, errors.New("queue is empty")
    }
}
```

### 3. Memory Pool for Frequent Allocations

```go
var nodePool = sync.Pool{
    New: func() interface{} {
        return &Node{}
    },
}

func getNode() *Node {
    return nodePool.Get().(*Node)
}

func putNode(n *Node) {
    n.data = nil
    n.next = nil
    nodePool.Put(n)
}
```

### 4. Zero-Allocation Techniques

```go
// Reuse slice to avoid allocations
func (s *Stack[T]) PopN(n int, buffer []T) []T {
    if cap(buffer) < n {
        buffer = make([]T, n)
    }
    buffer = buffer[:0] // Reset length, keep capacity
    
    for i := 0; i < n && !s.IsEmpty(); i++ {
        if item, ok := s.Pop(); ok {
            buffer = append(buffer, item)
        }
    }
    return buffer
}
```

## 🧪 Testing Patterns

### 1. Table-Driven Tests

```go
func TestStack(t *testing.T) {
    tests := []struct {
        name        string
        operations  []string
        values      []int
        expected    []int
        shouldError []bool
    }{
        {
            name:       "basic operations",
            operations: []string{"push", "push", "pop", "push", "pop", "pop"},
            values:     []int{1, 2, 0, 3, 0, 0},
            expected:   []int{0, 0, 2, 0, 3, 1},
            shouldError: []bool{false, false, false, false, false, false},
        },
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // Test implementation
        })
    }
}
```

### 2. Benchmark Comparisons

```go
func BenchmarkStackImplementations(b *testing.B) {
    implementations := map[string]func() Stack{
        "ArrayStack":  func() Stack { return NewArrayStack() },
        "LinkedStack": func() Stack { return NewLinkedStack() },
        "GenericStack": func() Stack { return NewGenericStack[int]() },
    }
    
    for name, factory := range implementations {
        b.Run(name, func(b *testing.B) {
            stack := factory()
            b.ResetTimer()
            
            for i := 0; i < b.N; i++ {
                stack.Push(i)
            }
        })
    }
}
```

## 🎯 Common Problem Patterns

### 1. Two Pointers

```go
func reverseArray(arr []int) {
    left, right := 0, len(arr)-1
    for left < right {
        arr[left], arr[right] = arr[right], arr[left]
        left++
        right--
    }
}
```

### 2. Sliding Window

```go
func maxSumSubarray(arr []int, k int) int {
    if len(arr) < k {
        return 0
    }
    
    // Calculate sum of first window
    windowSum := 0
    for i := 0; i < k; i++ {
        windowSum += arr[i]
    }
    
    maxSum := windowSum
    // Slide the window
    for i := k; i < len(arr); i++ {
        windowSum = windowSum - arr[i-k] + arr[i]
        if windowSum > maxSum {
            maxSum = windowSum
        }
    }
    
    return maxSum
}
```

### 3. Stack for Expression Evaluation

```go
func isBalanced(s string) bool {
    stack := NewGenericStack[rune]()
    pairs := map[rune]rune{')': '(', '}': '{', ']': '['}
    
    for _, char := range s {
        if char == '(' || char == '{' || char == '[' {
            stack.Push(char)
        } else if char == ')' || char == '}' || char == ']' {
            if stack.IsEmpty() {
                return false
            }
            if top, _ := stack.Pop(); top != pairs[char] {
                return false
            }
        }
    }
    
    return stack.IsEmpty()
}
```

---
*This reference covers the most common patterns you'll use in DSA with Go. Keep it handy for quick lookups!*
