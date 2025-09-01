# Practice Problems - Stacks & Arrays

## 🎯 Problem Categories

### Level 1: Fundamentals (Complete 5/8)

#### 1. Valid Parentheses ✅ COMPLETED
**Problem**: Check if parentheses are properly balanced
**Skills**: Stack operations, string processing
**Implementation**: See `stack.go` - `BalancedParentheses()`

#### 2. Min Stack
**Problem**: Design a stack that supports push, pop, top, and retrieving minimum element in O(1)
```go
type MinStack struct {
    // TODO: Implement this
}

func (ms *MinStack) Push(val int) {
    // TODO
}

func (ms *MinStack) Pop() {
    // TODO  
}

func (ms *MinStack) Top() int {
    // TODO
}

func (ms *MinStack) GetMin() int {
    // TODO
}
```

#### 3. Evaluate Reverse Polish Notation ✅ COMPLETED  
**Problem**: Evaluate postfix expressions
**Skills**: Stack operations, arithmetic parsing
**Implementation**: See `stack.go` - `EvaluatePostfix()`

#### 4. Daily Temperatures
**Problem**: Given daily temperatures, find how many days you have to wait for warmer temperature
```go
func dailyTemperatures(temperatures []int) []int {
    // TODO: Return array where answer[i] is the number of days 
    // you have to wait after the ith day to get a warmer temperature
}
```

#### 5. Remove Duplicates from Sorted Array
**Problem**: Remove duplicates in-place from sorted array
```go
func removeDuplicates(nums []int) int {
    // TODO: Modify nums in-place and return the new length
}
```

#### 6. Two Sum
**Problem**: Find indices of two numbers that add up to target
```go
func twoSum(nums []int, target int) []int {
    // TODO: Return indices of the two numbers
}
```

#### 7. Container With Most Water
**Problem**: Find two lines that together with x-axis form container with most water
```go
func maxArea(height []int) int {
    // TODO: Use two pointers technique
}
```

#### 8. Sliding Window Maximum
**Problem**: Find maximum in each sliding window of size k
```go
func maxSlidingWindow(nums []int, k int) []int {
    // TODO: Use deque for O(n) solution
}
```

### Level 2: Intermediate (Complete 3/6)

#### 1. Largest Rectangle in Histogram
**Problem**: Find area of largest rectangle in histogram
```go
func largestRectangleArea(heights []int) int {
    // TODO: Use stack to track indices
    // Hint: Process each bar and calculate area when stack pops
}
```

#### 2. Next Greater Element
**Problem**: Find next greater element for each element in array
```go
func nextGreaterElement(nums []int) []int {
    // TODO: Use stack to keep track of elements waiting for next greater
}
```

#### 3. Decode String
**Problem**: Decode string like "3[a2[c]]" → "accaccacc"
```go
func decodeString(s string) string {
    // TODO: Use stack to handle nested brackets
}
```

#### 4. Implement Queue using Stacks
**Problem**: Implement queue operations using only stacks
```go
type MyQueue struct {
    // TODO: Use two stacks
}

func (q *MyQueue) Push(x int) {
    // TODO
}

func (q *MyQueue) Pop() int {
    // TODO
}

func (q *MyQueue) Peek() int {
    // TODO
}

func (q *MyQueue) Empty() bool {
    // TODO
}
```

#### 5. Trapping Rain Water
**Problem**: Calculate trapped rainwater in elevation map
```go
func trap(height []int) int {
    // TODO: Use two pointers or stack approach
}
```

#### 6. Asteroid Collision
**Problem**: Simulate asteroid collisions
```go
func asteroidCollision(asteroids []int) []int {
    // TODO: Use stack to simulate collisions
}
```

### Level 3: Advanced (Complete 2/4)

#### 1. Basic Calculator
**Problem**: Implement calculator for expressions with +, -, (, )
```go
func calculate(s string) int {
    // TODO: Use stack to handle parentheses and operations
}
```

#### 2. Longest Valid Parentheses
**Problem**: Find length of longest valid parentheses substring
```go
func longestValidParentheses(s string) int {
    // TODO: Use stack to track indices
}
```

#### 3. Maximum Frequency Stack
**Problem**: Implement stack where pop returns most frequent element
```go
type FreqStack struct {
    // TODO: Track frequency and maintain stacks for each frequency
}

func (fs *FreqStack) Push(val int) {
    // TODO
}

func (fs *FreqStack) Pop() int {
    // TODO: Return most frequent element
}
```

#### 4. Sliding Window Median
**Problem**: Find median in each sliding window
```go
func medianSlidingWindow(nums []int, k int) []float64 {
    // TODO: Use two heaps or balanced data structure
}
```

## 🧪 Testing Template

For each problem, use this testing pattern:

```go
func TestProblemName(t *testing.T) {
    tests := []struct {
        name     string
        input    InputType
        expected OutputType
    }{
        {
            name:     "basic case",
            input:    InputType{...},
            expected: OutputType{...},
        },
        {
            name:     "edge case - empty",
            input:    InputType{},
            expected: OutputType{},
        },
        {
            name:     "edge case - single element",
            input:    InputType{...},
            expected: OutputType{...},
        },
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := problemFunction(tt.input)
            if !reflect.DeepEqual(result, tt.expected) {
                t.Errorf("got %v, want %v", result, tt.expected)
            }
        })
    }
}
```

## 📊 Progress Tracking

### Week 1-2: Stack Fundamentals ✅
- [x] Basic stack operations
- [x] Stack applications (parentheses, postfix)
- [x] Multiple implementations (array, linked, generic)

### Week 3: Array Manipulation 🚧
- [ ] Two pointers technique
- [ ] Sliding window problems  
- [ ] In-place modifications
- [ ] Search and sorting

### Week 4: Advanced Stack Problems 📋
- [ ] Calculator implementation
- [ ] Histogram problems
- [ ] String processing with stacks
- [ ] Frequency-based problems

## 🎯 Learning Objectives

### Technical Skills:
- Master stack-based problem solving
- Understand when to use different data structures
- Implement efficient algorithms
- Handle edge cases properly

### Problem-Solving Patterns:
- **Monotonic Stack**: For next/previous greater elements
- **Two Stacks**: For queue implementation or expression evaluation
- **Stack + HashMap**: For frequency tracking
- **Stack + String**: For parsing and validation

## 💡 Hints & Strategies

### Stack Problems:
1. **Parentheses/Brackets**: Always use stack for nested structures
2. **Expression Evaluation**: Stack for operators, consider precedence
3. **Histogram/Area**: Stack to track increasing sequences
4. **String Parsing**: Stack for handling nested or recursive patterns

### Array Problems:
1. **Two Pointers**: For sorted arrays or when you need O(1) space
2. **Sliding Window**: For subarray problems with size constraint
3. **Hash Map**: For O(1) lookups and frequency counting
4. **In-place**: When space complexity matters

## 📈 Difficulty Progression

```
Fundamentals → Applications → Optimizations → System Design

Valid Parentheses → Min Stack → Calculator → Distributed Calculator
Two Sum → Container Water → Trapping Water → Rate Limiter
Basic Array → Sliding Window → Advanced Window → Cache Implementation
```

## 🔄 Review Schedule

- **Daily**: Solve 1-2 problems, review solutions
- **Weekly**: Benchmark implementations, optimize bottlenecks  
- **Bi-weekly**: Implement real-world applications
- **Monthly**: System design using learned concepts

---
*Complete problems in order, focus on understanding patterns, not just solutions!*
