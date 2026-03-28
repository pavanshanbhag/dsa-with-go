# Complete Golang Learning Guide

*A step-by-step guide to learning the Go programming language*

**Use this guide if you're new to Go.** The DSA content in this repo (01–05) assumes basic Go syntax, types, and error handling. For the **DSA learning path** (what to study in what order), see [LEARNING_PATH.md](LEARNING_PATH.md).

---

## 📋 Table of Contents

1. [🎯 Learning Path Overview](#-learning-path-overview)
2. [📝 Prerequisites](#-prerequisites)
3. [🏗️ Development Environment Setup](#️-development-environment-setup)
4. [🔰 Phase 1: Go Fundamentals](#-phase-1-go-fundamentals)
5. [🏗️ Phase 2: Core Programming Concepts](#️-phase-2-core-programming-concepts)
6. [🚀 Phase 3: Intermediate Go Features](#-phase-3-intermediate-go-features)
7. [⚡ Phase 4: Advanced Go Programming](#-phase-4-advanced-go-programming)
8. [🌐 Phase 5: Web Development & APIs](#-phase-5-web-development--apis)
9. [🔧 Phase 6: Production-Ready Skills](#-phase-6-production-ready-skills)
10. [📚 Additional Resources](#-additional-resources)
11. [🎯 Practice Projects](#-practice-projects)
12. [📈 Assessment & Next Steps](#-assessment--next-steps)

---

## 🎯 Learning Path Overview

This guide is designed to take you from a complete Go beginner to an advanced Go developer capable of building production-ready applications. The estimated timeline is **8-12 weeks** with consistent daily practice.

### Learning Phases Timeline

| Phase | Duration | Focus Area | Skills Gained |
|-------|----------|------------|---------------|
| **Phase 1** | Week 1-2 | Go Fundamentals | Basic syntax, data types, control flow |
| **Phase 2** | Week 2-4 | Core Programming | Functions, structs, methods, interfaces |
| **Phase 3** | Week 4-6 | Intermediate Features | Goroutines, channels, error handling |
| **Phase 4** | Week 6-8 | Advanced Programming | Reflection, generics, performance optimization |
| **Phase 5** | Week 8-10 | Web Development | HTTP servers, REST APIs, databases |
| **Phase 6** | Week 10-12 | Production Skills | Testing, deployment, monitoring |

---

## 📝 Prerequisites

### Required Knowledge

- Basic programming concepts (variables, loops, functions)
- Command-line/terminal usage
- Basic understanding of version control (Git)

### Recommended Background

- Experience with any programming language (Python, Java, C++, JavaScript, etc.)
- Understanding of basic computer science concepts

---

## 🏗️ Development Environment Setup

### 1. Install Go

```bash
# Check if Go is already installed
go version

# If not installed, download from https://go.dev/dl/
# Or install via package manager:

# macOS (using Homebrew)
brew install go

# Ubuntu/Debian
sudo apt update
sudo apt install golang-go

# Windows (using Chocolatey)
choco install golang
```

### 2. Configure Environment

```bash
# Add to your shell profile (.bashrc, .zshrc, etc.)
export GOPATH=$HOME/go
export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin

# Verify installation
go version
go env GOPATH
go env GOROOT
```

### 3. IDE Setup

**Recommended IDEs:**

- **VS Code** with Go extension (recommended for beginners)
- **GoLand** by JetBrains (professional)
- **Vim/Neovim** with go plugins (for advanced users)

**VS Code Extensions:**

```bash
# Install Go extension
code --install-extension golang.go
```

### 4. Initialize Your Learning Workspace

```bash
mkdir ~/go-learning
cd ~/go-learning
go mod init go-learning
```

---

## 🔰 Phase 1: Go Fundamentals

*Duration: 2 weeks | Goal: Master basic Go syntax and concepts*

### Week 1: Getting Started

#### Day 1-2: Hello World & Basic Syntax

**Concepts to Learn:**

- Go program structure
- Package system
- Basic I/O operations
- Comments and code organization

**Practice Exercise:**

```go
// hello.go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
    fmt.Printf("Welcome to Go programming!\n")
}
```

**Tasks:**

- [ ] Create your first Go program
- [ ] Understand package main and import statements
- [ ] Practice with fmt package functions
- [ ] Learn go run vs go build

#### Day 3-4: Variables and Data Types

**Concepts to Learn:**

- Variable declarations (var, :=)
- Basic data types (int, float, string, bool)
- Type inference
- Constants
- Zero values

**Practice Exercise:**

```go
package main

import "fmt"

func main() {
    // Variable declarations
    var name string = "Go"
    age := 13
    var isAwesome bool = true
    const pi = 3.14159
    
    fmt.Printf("Language: %s, Age: %d, Awesome: %t, Pi: %.2f\n", 
               name, age, isAwesome, pi)
}
```

**Tasks:**

- [ ] Practice different variable declaration styles
- [ ] Experiment with all basic data types
- [ ] Understand type conversion
- [ ] Work with constants

#### Day 5-7: Control Flow

**Concepts to Learn:**

- if/else statements
- for loops (various forms)
- switch statements
- defer statement

**Practice Exercise:**

```go
package main

import "fmt"

func main() {
    // if-else
    score := 85
    if score >= 90 {
        fmt.Println("Grade: A")
    } else if score >= 80 {
        fmt.Println("Grade: B")
    } else {
        fmt.Println("Grade: C")
    }
    
    // for loop variations
    for i := 0; i < 5; i++ {
        fmt.Printf("Iteration %d\n", i)
    }
    
    // switch statement
    day := "Monday"
    switch day {
    case "Monday":
        fmt.Println("Start of the work week")
    case "Friday":
        fmt.Println("TGIF!")
    default:
        fmt.Println("Regular day")
    }
}
```

**Tasks:**

- [ ] Build a simple calculator using switch
- [ ] Create number guessing game with loops
- [ ] Practice with defer statement
- [ ] Understand break and continue

### Week 2: Data Structures Basics

#### Day 8-10: Arrays and Slices

**Concepts to Learn:**

- Array declaration and initialization
- Slice creation and manipulation
- Slice operations (append, copy, len, cap)
- Multi-dimensional arrays

**Practice Exercise:**

```go
package main

import "fmt"

func main() {
    // Arrays
    var numbers [5]int = [5]int{1, 2, 3, 4, 5}
    
    // Slices
    fruits := []string{"apple", "banana", "orange"}
    fruits = append(fruits, "grape")
    
    fmt.Println("Numbers:", numbers)
    fmt.Println("Fruits:", fruits)
    fmt.Printf("Length: %d, Capacity: %d\n", len(fruits), cap(fruits))
}
```

**Tasks:**

- [ ] Implement array sorting without built-in functions
- [ ] Practice slice operations extensively
- [ ] Build a dynamic list management system
- [ ] Understand memory allocation with slices

#### Day 11-14: Maps and Strings

**Concepts to Learn:**

- Map creation and operations
- String manipulation
- Runes and Unicode
- String formatting

**Practice Exercise:**

```go
package main

import (
    "fmt"
    "strings"
)

func main() {
    // Maps
    ages := make(map[string]int)
    ages["Alice"] = 30
    ages["Bob"] = 25
    
    // String operations
    text := "Hello, Go Programming!"
    fmt.Println("Uppercase:", strings.ToUpper(text))
    fmt.Println("Contains 'Go':", strings.Contains(text, "Go"))
    
    // Iterate over map
    for name, age := range ages {
        fmt.Printf("%s is %d years old\n", name, age)
    }
}
```

**Tasks:**

- [ ] Build a word frequency counter
- [ ] Create a phone book application using maps
- [ ] Practice string manipulation functions
- [ ] Work with runes for Unicode text

**📊 Phase 1 Assessment:**

- [ ] Build a student grade management system
- [ ] Create a simple inventory tracker
- [ ] Implement basic text processing utilities

---

## 🏗️ Phase 2: Core Programming Concepts

*Duration: 2 weeks | Goal: Master functions, structs, and methods*

### Week 3: Functions and Error Handling

#### Day 15-17: Functions Deep Dive

**Concepts to Learn:**

- Function declaration and parameters
- Multiple return values
- Variadic functions
- Anonymous functions and closures
- Function as values

**Practice Exercise:**

```go
package main

import "fmt"

// Multiple return values
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero")
    }
    return a / b, nil
}

// Variadic function
func sum(numbers ...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}

func main() {
    result, err := divide(10, 3)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Printf("Result: %.2f\n", result)
    }
    
    fmt.Println("Sum:", sum(1, 2, 3, 4, 5))
}
```

**Tasks:**

- [ ] Implement calculator with proper error handling
- [ ] Create utility functions for common operations
- [ ] Practice with higher-order functions
- [ ] Build function composition examples

#### Day 18-21: Structs and Methods

**Concepts to Learn:**

- Struct definition and initialization
- Embedded structs
- Methods and receivers
- Pointer vs value receivers
- Struct tags

**Practice Exercise:**

```go
package main

import "fmt"

// Define struct
type Person struct {
    Name    string
    Age     int
    Email   string
}

// Method with value receiver
func (p Person) Greet() string {
    return fmt.Sprintf("Hello, I'm %s", p.Name)
}

// Method with pointer receiver
func (p *Person) HaveBirthday() {
    p.Age++
}

func main() {
    person := Person{
        Name:  "Alice",
        Age:   30,
        Email: "alice@example.com",
    }
    
    fmt.Println(person.Greet())
    person.HaveBirthday()
    fmt.Printf("Age after birthday: %d\n", person.Age)
}
```

**Tasks:**

- [ ] Design a library management system with structs
- [ ] Implement geometric shapes with methods
- [ ] Create a banking account system
- [ ] Practice with embedded structs

### Week 4: Interfaces and Packages

#### Day 22-24: Interfaces

**Concepts to Learn:**

- Interface definition and implementation
- Empty interface
- Type assertions
- Interface composition
- Common Go interfaces (io.Reader, io.Writer, etc.)

**Practice Exercise:**

```go
package main

import "fmt"

// Define interface
type Shape interface {
    Area() float64
    Perimeter() float64
}

type Rectangle struct {
    Width, Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
    return 2 * (r.Width + r.Height)
}

func PrintShapeInfo(s Shape) {
    fmt.Printf("Area: %.2f, Perimeter: %.2f\n", s.Area(), s.Perimeter())
}

func main() {
    rect := Rectangle{Width: 5, Height: 3}
    PrintShapeInfo(rect)
}
```

**Tasks:**

- [ ] Implement multiple shapes with Shape interface
- [ ] Create a plugin system using interfaces
- [ ] Practice with io.Reader and io.Writer
- [ ] Build a sorting system with interfaces

#### Day 25-28: Packages and Modules

**Concepts to Learn:**

- Package creation and organization
- Go modules (go.mod, go.sum)
- Importing packages
- Package visibility (exported/unexported)
- Third-party packages

**Practice Exercise:**

```go
// mathutils/calculator.go
package mathutils

// Exported function (starts with capital letter)
func Add(a, b int) int {
    return a + b
}

// unexported function (starts with lowercase)
func multiply(a, b int) int {
    return a * b
}

// Exported function using unexported function
func Square(n int) int {
    return multiply(n, n)
}
```

**Tasks:**

- [ ] Create your own utility package
- [ ] Use popular third-party packages (chi, echo, etc.)
- [ ] Organize code into logical packages
- [ ] Practice with go mod commands

**📊 Phase 2 Assessment:**

- [ ] Build a complete task management CLI application
- [ ] Create a file processing utility with interfaces
- [ ] Design a modular calculator package

---

## 🚀 Phase 3: Intermediate Go Features

*Duration: 2 weeks | Goal: Master concurrency and advanced error handling*

### Week 5: Concurrency Fundamentals

#### Day 29-31: Goroutines

**Concepts to Learn:**

- Goroutine creation and lifecycle
- Go scheduler basics
- Goroutine synchronization
- Race conditions and data races

**Practice Exercise:**

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func worker(id int, wg *sync.WaitGroup) {
    defer wg.Done()
    fmt.Printf("Worker %d starting\n", id)
    time.Sleep(time.Second)
    fmt.Printf("Worker %d done\n", id)
}

func main() {
    var wg sync.WaitGroup
    
    for i := 1; i <= 3; i++ {
        wg.Add(1)
        go worker(i, &wg)
    }
    
    wg.Wait()
    fmt.Println("All workers completed")
}
```

**Tasks:**

- [ ] Build a concurrent file processor
- [ ] Create a web scraper using goroutines
- [ ] Implement a worker pool pattern
- [ ] Practice with sync.WaitGroup and sync.Mutex

#### Day 32-35: Channels

**Concepts to Learn:**

- Channel creation and operations
- Buffered vs unbuffered channels
- Channel direction (send-only, receive-only)
- Select statement
- Channel patterns (fan-in, fan-out, pipeline)

**Practice Exercise:**

```go
package main

import (
    "fmt"
    "time"
)

func producer(ch chan<- int) {
    for i := 1; i <= 5; i++ {
        ch <- i
        time.Sleep(time.Millisecond * 500)
    }
    close(ch)
}

func consumer(ch <-chan int) {
    for value := range ch {
        fmt.Printf("Received: %d\n", value)
    }
}

func main() {
    ch := make(chan int)
    
    go producer(ch)
    consumer(ch)
    
    fmt.Println("Done")
}
```

**Tasks:**

- [ ] Implement a pub-sub system with channels
- [ ] Create a pipeline processing system
- [ ] Build a chat server using channels
- [ ] Practice with select statement and timeouts

### Week 6: Advanced Error Handling and Context

#### Day 36-38: Error Handling Patterns

**Concepts to Learn:**

- Custom error types
- Error wrapping and unwrapping
- Panic and recover
- Error handling best practices

**Practice Exercise:**

```go
package main

import (
    "errors"
    "fmt"
)

// Custom error type
type ValidationError struct {
    Field string
    Value interface{}
    Message string
}

func (e ValidationError) Error() string {
    return fmt.Sprintf("validation failed for %s: %s", e.Field, e.Message)
}

func validateAge(age int) error {
    if age < 0 || age > 150 {
        return ValidationError{
            Field:   "age",
            Value:   age,
            Message: "age must be between 0 and 150",
        }
    }
    return nil
}

func main() {
    if err := validateAge(-5); err != nil {
        var valErr ValidationError
        if errors.As(err, &valErr) {
            fmt.Printf("Validation error: %s\n", valErr.Error())
        }
    }
}
```

**Tasks:**

- [ ] Create a robust input validation system
- [ ] Implement graceful error recovery
- [ ] Build error logging utilities
- [ ] Practice with error wrapping

#### Day 39-42: Context Package

**Concepts to Learn:**

- Context creation and propagation
- Cancellation and timeouts
- Context values
- Context best practices

**Practice Exercise:**

```go
package main

import (
    "context"
    "fmt"
    "time"
)

func longRunningTask(ctx context.Context) error {
    select {
    case <-time.After(5 * time.Second):
        fmt.Println("Task completed")
        return nil
    case <-ctx.Done():
        return ctx.Err()
    }
}

func main() {
    ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
    defer cancel()
    
    if err := longRunningTask(ctx); err != nil {
        fmt.Printf("Task failed: %v\n", err)
    }
}
```

**Tasks:**

- [ ] Build an HTTP client with proper context handling
- [ ] Implement cancellable operations
- [ ] Create context-aware database operations
- [ ] Practice with context propagation in web handlers

**📊 Phase 3 Assessment:**

- [ ] Build a concurrent web crawler
- [ ] Create a job processing system with workers
- [ ] Implement a robust HTTP client library

---

## ⚡ Phase 4: Advanced Go Programming

*Duration: 2 weeks | Goal: Master advanced features and optimization*

### Week 7: Generics and Reflection

#### Day 43-45: Go Generics (Go 1.18+, 1.24+ recommended)

**Concepts to Learn:**

- Generic functions and types
- Type constraints
- Type inference
- Generic interfaces

**Practice Exercise:**

```go
package main

import (
    "cmp"
    "fmt"
)

// Generic function using cmp.Ordered (Go 1.24+)
func Max[T cmp.Ordered](a, b T) T {
    if a > b {
        return a
    }
    return b
}

// Generic type
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

func main() {
    fmt.Println("Max of 5 and 10:", Max(5, 10))
    fmt.Println("Max of 'b' and 'a':", Max('b', 'a'))
    
    stack := Stack[int]{}
    stack.Push(1)
    stack.Push(2)
    
    if item, ok := stack.Pop(); ok {
        fmt.Println("Popped:", item)
    }
}
```

**Tasks:**

- [ ] Implement generic data structures (list, map, set)
- [ ] Create utility functions with type constraints
- [ ] Build a generic sorting library
- [ ] Practice with complex type constraints

#### Day 46-49: Reflection and Performance

**Concepts to Learn:**

- Reflection basics (reflect package)
- Type and value inspection
- Dynamic method calls
- Performance considerations
- Memory profiling

**Practice Exercise:**

```go
package main

import (
    "fmt"
    "reflect"
)

type User struct {
    Name  string `json:"name"`
    Age   int    `json:"age"`
    Email string `json:"email"`
}

func printFields(v interface{}) {
    val := reflect.ValueOf(v)
    typ := reflect.TypeOf(v)
    
    for i := 0; i < val.NumField(); i++ {
        field := val.Field(i)
        fieldType := typ.Field(i)
        tag := fieldType.Tag.Get("json")
        
        fmt.Printf("Field: %s, Type: %s, Value: %v, Tag: %s\n",
            fieldType.Name, field.Type(), field.Interface(), tag)
    }
}

func main() {
    user := User{Name: "Alice", Age: 30, Email: "alice@example.com"}
    printFields(user)
}
```

**Tasks:**

- [ ] Build a simple ORM using reflection
- [ ] Create a configuration parser with struct tags
- [ ] Implement a generic validator
- [ ] Profile and optimize performance-critical code

### Week 8: Testing and Documentation

#### Day 50-52: Advanced Testing

**Concepts to Learn:**

- Unit testing best practices
- Table-driven tests
- Benchmarking
- Test coverage
- Mocking and dependency injection

**Practice Exercise:**

```go
// calculator.go
package main

import "errors"

func Add(a, b int) int {
    return a + b
}

func Divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

// calculator_test.go
package main

import (
    "testing"
)

func TestAdd(t *testing.T) {
    tests := []struct {
        name     string
        a, b     int
        expected int
    }{
        {"positive numbers", 5, 3, 8},
        {"negative numbers", -2, -3, -5},
        {"zero", 0, 5, 5},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := Add(tt.a, tt.b)
            if result != tt.expected {
                t.Errorf("Add(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
            }
        })
    }
}

func BenchmarkAdd(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Add(100, 200)
    }
}
```

**Tasks:**

- [ ] Write comprehensive tests for all previous projects
- [ ] Implement benchmark tests
- [ ] Create integration tests
- [ ] Practice with test mocks

#### Day 53-56: Documentation and Code Quality

**Concepts to Learn:**

- Godoc documentation
- Code organization best practices
- Linting and formatting (gofmt, golint, go vet)
- Static analysis tools

**Practice Exercise:**

```go
// Package calculator provides basic arithmetic operations.
package calculator

import "errors"

// Add returns the sum of two integers.
//
// Example:
//   result := Add(5, 3) // result is 8
func Add(a, b int) int {
    return a + b
}

// Divide performs division of two float64 numbers.
// It returns an error if the divisor is zero.
//
// Example:
//   result, err := Divide(10.0, 2.0)
//   if err != nil {
//       log.Fatal(err)
//   }
//   fmt.Println(result) // Output: 5.0
func Divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}
```

**Tasks:**

- [ ] Document all packages with godoc
- [ ] Set up pre-commit hooks for code quality
- [ ] Create comprehensive README files
- [ ] Practice with static analysis tools

**📊 Phase 4 Assessment:**

- [ ] Build a generic data processing library
- [ ] Create a well-tested and documented package
- [ ] Implement performance-optimized algorithms

---

## 🌐 Phase 5: Web Development & APIs

*Duration: 2 weeks | Goal: Build web applications and REST APIs*

### Week 9: HTTP and REST APIs

#### Day 57-59: HTTP Server Basics

**Concepts to Learn:**

- net/http package
- HTTP handlers and middleware
- Routing patterns
- Request and response handling
- Static file serving

**Practice Exercise:**

```go
package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
)

type User struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

var users = []User{
    {ID: 1, Name: "Alice"},
    {ID: 2, Name: "Bob"},
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(users)
}

func main() {
    http.HandleFunc("/users", usersHandler)
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Welcome to Go Web Server!")
    })
    
    fmt.Println("Server starting on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

**Tasks:**

- [ ] Build a complete REST API for a todo application
- [ ] Implement middleware for logging and authentication
- [ ] Create file upload/download endpoints
- [ ] Practice with different HTTP methods

#### Day 60-63: Advanced Web Features

**Concepts to Learn:**

- Third-party routers (e.g. chi, echo) or std lib net/http ServeMux
- Database integration (SQL, MongoDB)
- Authentication and authorization
- WebSockets
- Template rendering

**Practice Exercise:**

```go
package main

import (
    "database/sql"
    "encoding/json"
    "log"
    "net/http"

    "github.com/go-chi/chi/v5"
    _ "github.com/lib/pq"
)

type Article struct {
    ID    int    `json:"id" db:"id"`
    Title string `json:"title" db:"title"`
    Body  string `json:"body" db:"body"`
}

type Server struct {
    db *sql.DB
}

func (s *Server) getArticles(w http.ResponseWriter, r *http.Request) {
    rows, err := s.db.Query("SELECT id, title, body FROM articles")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()
    
    var articles []Article
    for rows.Next() {
        var article Article
        err := rows.Scan(&article.ID, &article.Title, &article.Body)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        articles = append(articles, article)
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(articles)
}

func main() {
    // Database connection setup
    db, err := sql.Open("postgres", "postgresql://user:password@localhost/dbname?sslmode=disable")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
    
    server := &Server{db: db}
    
    r := chi.NewRouter()
    r.Get("/articles", server.getArticles)
    
    http.ListenAndServe(":8080", r)
}
```

**Tasks:**

- [ ] Build a complete blog API with database
- [ ] Implement JWT authentication
- [ ] Create real-time features with WebSockets
- [ ] Practice with different databases

### Week 10: Microservices and APIs

#### Day 64-66: Microservices Architecture

**Concepts to Learn:**

- Microservices design patterns
- Service communication (HTTP, gRPC)
- Configuration management
- Health checks and monitoring

**Practice Exercise:**

```go
// user-service/main.go
package main

import (
    "encoding/json"
    "net/http"
)

type User struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

type UserService struct {
    users map[int]User
}

func (us *UserService) getUserHandler(w http.ResponseWriter, r *http.Request) {
    // Implementation here
}

func (us *UserService) healthHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"status": "healthy"})
}

func main() {
    us := &UserService{
        users: make(map[int]User),
    }
    
    http.HandleFunc("/users/", us.getUserHandler)
    http.HandleFunc("/health", us.healthHandler)
    
    http.ListenAndServe(":8081", nil)
}
```

**Tasks:**

- [ ] Build multiple communicating microservices
- [ ] Implement service discovery
- [ ] Create API gateway pattern
- [ ] Practice with gRPC

#### Day 67-70: API Documentation and Testing

**Concepts to Learn:**

- OpenAPI/Swagger documentation
- API testing strategies
- Load testing
- API versioning

**Tasks:**

- [ ] Document APIs with Swagger
- [ ] Implement comprehensive API tests
- [ ] Create load testing scenarios
- [ ] Practice API versioning strategies

**📊 Phase 5 Assessment:**

- [ ] Build a complete e-commerce microservices system
- [ ] Create a real-time chat application
- [ ] Implement a RESTful API with full CRUD operations

---

## 🔧 Phase 6: Production-Ready Skills

*Duration: 2 weeks | Goal: Deploy and maintain production applications*

### Week 11: Deployment and DevOps

#### Day 71-73: Containerization

**Concepts to Learn:**

- Docker containerization
- Multi-stage builds
- Docker Compose
- Container optimization

**Practice Exercise:**

```dockerfile
# Dockerfile
FROM golang:1.24-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]
```

**Tasks:**

- [ ] Containerize all previous projects
- [ ] Create Docker Compose for multi-service applications
- [ ] Optimize container sizes
- [ ] Practice with container orchestration

#### Day 74-77: Monitoring and Logging

**Concepts to Learn:**

- Structured logging
- Metrics collection (Prometheus)
- Distributed tracing
- Health monitoring

**Practice Exercise:**

```go
package main

import (
    "log/slog"
    "net/http"
    "time"
    
    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
    httpRequestsTotal = prometheus.NewCounterVec(
        prometheus.CounterOpts{
            Name: "http_requests_total",
            Help: "Total number of HTTP requests",
        },
        []string{"method", "endpoint"},
    )
    
    httpRequestDuration = prometheus.NewHistogramVec(
        prometheus.HistogramOpts{
            Name: "http_request_duration_seconds",
            Help: "HTTP request duration in seconds",
        },
        []string{"method", "endpoint"},
    )
)

func init() {
    prometheus.MustRegister(httpRequestsTotal)
    prometheus.MustRegister(httpRequestDuration)
}

func instrumentHandler(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        
        slog.Info("Request started",
            "method", r.Method,
            "path", r.URL.Path,
            "remote_addr", r.RemoteAddr,
        )
        
        next(w, r)
        
        duration := time.Since(start).Seconds()
        httpRequestsTotal.WithLabelValues(r.Method, r.URL.Path).Inc()
        httpRequestDuration.WithLabelValues(r.Method, r.URL.Path).Observe(duration)
        
        slog.Info("Request completed",
            "method", r.Method,
            "path", r.URL.Path,
            "duration_seconds", duration,
        )
    }
}

func main() {
    http.HandleFunc("/", instrumentHandler(func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello, monitored world!"))
    }))
    
    http.Handle("/metrics", promhttp.Handler())
    
    slog.Info("Server starting on :8080")
    http.ListenAndServe(":8080", nil)
}
```

**Tasks:**

- [ ] Implement comprehensive logging in applications
- [ ] Set up Prometheus metrics collection
- [ ] Create monitoring dashboards
- [ ] Practice with distributed tracing

### Week 12: Security and Performance

#### Day 78-80: Security Best Practices

**Concepts to Learn:**

- Input validation and sanitization
- HTTPS and TLS configuration
- Authentication and authorization patterns
- Security scanning and vulnerability assessment

**Practice Exercise:**

```go
package main

import (
    "crypto/tls"
    "net/http"
    "time"
)

func secureMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Security headers
        w.Header().Set("X-Content-Type-Options", "nosniff")
        w.Header().Set("X-Frame-Options", "DENY")
        w.Header().Set("X-XSS-Protection", "1; mode=block")
        w.Header().Set("Strict-Transport-Security", "max-age=31536000")
        
        next.ServeHTTP(w, r)
    })
}

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Secure server"))
    })
    
    server := &http.Server{
        Addr:         ":8443",
        Handler:      secureMiddleware(mux),
        TLSConfig:    &tls.Config{MinVersion: tls.VersionTLS12},
        ReadTimeout:  10 * time.Second,
        WriteTimeout: 10 * time.Second,
        IdleTimeout:  120 * time.Second,
    }
    
    server.ListenAndServeTLS("cert.pem", "key.pem")
}
```

**Tasks:**

- [ ] Implement secure authentication systems
- [ ] Set up HTTPS with proper TLS configuration
- [ ] Create input validation libraries
- [ ] Practice with security scanning tools

#### Day 81-84: Performance Optimization

**Concepts to Learn:**

- Performance profiling (pprof)
- Memory optimization
- CPU optimization
- Database query optimization
- Caching strategies

**Practice Exercise:**

```go
package main

import (
    "fmt"
    "runtime"
    "time"
    _ "net/http/pprof"
    "net/http"
)

// Optimized function with memory pool
type Pool struct {
    buffers chan []byte
}

func NewPool(size int, bufferSize int) *Pool {
    pool := &Pool{
        buffers: make(chan []byte, size),
    }
    
    for i := 0; i < size; i++ {
        pool.buffers <- make([]byte, 0, bufferSize)
    }
    
    return pool
}

func (p *Pool) Get() []byte {
    select {
    case buffer := <-p.buffers:
        return buffer[:0]
    default:
        return make([]byte, 0, 1024)
    }
}

func (p *Pool) Put(buffer []byte) {
    select {
    case p.buffers <- buffer:
    default:
    }
}

func main() {
    // Enable pprof endpoint
    go func() {
        http.ListenAndServe(":6060", nil)
    }()
    
    pool := NewPool(10, 1024)
    
    // Simulate work with memory optimization
    for i := 0; i < 1000; i++ {
        buffer := pool.Get()
        
        // Do work with buffer
        for j := 0; j < 100; j++ {
            buffer = append(buffer, byte(j))
        }
        
        pool.Put(buffer)
        
        if i%100 == 0 {
            runtime.GC()
            var m runtime.MemStats
            runtime.ReadMemStats(&m)
            fmt.Printf("Iteration %d: Alloc = %d KB", i, m.Alloc/1024)
        }
    }
}
```

**Tasks:**

- [ ] Profile and optimize existing applications
- [ ] Implement caching strategies
- [ ] Optimize database interactions
- [ ] Create performance benchmarks

**📊 Phase 6 Assessment:**

- [ ] Deploy a production-ready application
- [ ] Implement comprehensive monitoring
- [ ] Create security-hardened services
- [ ] Optimize application performance

---

## 📚 Additional Resources

### 📖 Recommended Books

1. **"The Go Programming Language"** by Alan Donovan and Brian Kernighan
2. **"Go in Action"** by William Kennedy, Brian Ketelsen, and Erik St. Martin
3. **"Concurrency in Go"** by Katherine Cox-Buday
4. **"Cloud Native Go"** by Matthew Titmus

### 🌐 Online Resources

- **Official Go Documentation**: <https://go.dev/doc/>
- **Go by Example**: <https://gobyexample.com/>
- **Effective Go**: <https://go.dev/doc/effective_go>
- **Go Blog**: <https://go.dev/blog>
- **A Tour of Go**: <https://go.dev/tour/>

### 🎥 Video Courses

- **Go Tutorial for Beginners** (TechWorld with Nana)
- **Go: The Complete Developer's Guide** (Udemy)
- **Building Modern Web Applications with Go** (Pluralsight)

### 🛠️ Development Tools

- **gofmt** / **go fmt**: Code formatting
- **go vet**: Static analysis
- **golangci-lint**: Linting (recommended; replaces deprecated golint)
- **staticcheck**: Additional static checks
- **pprof**: Performance profiling
- **delve**: Debugging

---

## 🎯 Practice Projects

### Beginner Projects

1. **CLI Calculator**: Basic arithmetic operations with error handling
2. **File Organizer**: Sort files by type and date
3. **Password Generator**: Configurable password creation tool
4. **Todo CLI**: Command-line todo list manager

### Intermediate Projects

1. **URL Shortener**: REST API with database persistence
2. **Chat Server**: Real-time messaging with WebSockets
3. **Log Parser**: Analyze and report on log files
4. **API Gateway**: Route requests to multiple backend services

### Advanced Projects

1. **Distributed Cache**: Redis-like caching system
2. **Container Orchestrator**: Simple Kubernetes-like system
3. **Monitoring System**: Metrics collection and alerting
4. **Microservices Platform**: Complete e-commerce system

---

## 📈 Assessment & Next Steps

### Self-Assessment Checklist

**Fundamentals** ✅

- [ ] Can write idiomatic Go code
- [ ] Understand Go's type system
- [ ] Comfortable with control flow
- [ ] Know standard library basics

**Core Programming** ✅

- [ ] Design effective structs and interfaces
- [ ] Write comprehensive tests
- [ ] Handle errors properly
- [ ] Organize code into packages

**Concurrency** ✅

- [ ] Use goroutines effectively
- [ ] Understand channel patterns
- [ ] Avoid race conditions
- [ ] Design concurrent systems

**Web Development** ✅

- [ ] Build REST APIs
- [ ] Handle HTTP requests/responses
- [ ] Integrate with databases
- [ ] Implement authentication

**Production Skills** ✅

- [ ] Deploy applications
- [ ] Monitor system health
- [ ] Optimize performance
- [ ] Secure applications

### Career Paths

1. **Backend Developer**: Focus on APIs, databases, and microservices
2. **DevOps Engineer**: Emphasis on deployment, monitoring, and infrastructure
3. **System Programmer**: Low-level programming and performance optimization
4. **Cloud Engineer**: Cloud-native applications and distributed systems

### Continuing Education

1. **Advanced Topics**:
   - Kubernetes and container orchestration
   - gRPC and protocol buffers
   - Database design and optimization
   - Distributed systems patterns

2. **Community Involvement**:
   - Contribute to open-source Go projects
   - Attend Go conferences and meetups
   - Write technical blog posts
   - Mentor other Go learners

3. **Certifications**:
   - Cloud provider certifications (AWS, GCP, Azure)
   - Kubernetes certifications (CKA, CKAD)
   - Security certifications

---

## 🎉 Conclusion

Congratulations on completing this comprehensive Go learning guide! You've built a solid foundation in Go programming and are ready to tackle real-world projects. Remember:

- **Practice consistently**: The key to mastering Go is regular practice
- **Build projects**: Apply your knowledge to real-world problems
- **Stay updated**: Go evolves rapidly, keep learning new features
- **Join the community**: Engage with other Go developers

**Next Steps:**

1. Choose a career path that interests you
2. Build a portfolio of Go projects
3. Contribute to open-source projects
4. Stay connected with the Go community

Happy coding! 🚀

---

*Last updated: February 2026*  
*Go version: 1.24+*
