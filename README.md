# DSA With Go: Complete Data Structures and Algorithms

A comprehensive, production-ready implementation of Data Structures and Algorithms in Go, designed for mastering DSA concepts.

## 🎯 Overview

This project provides a complete learning path for Data Structures and Algorithms in Go, featuring:

- **Comprehensive Coverage**: From basic complexity analysis to advanced graph algorithms
- **Production Quality**: Clean, efficient, and well-tested implementations
- **Educational Focus**: Clear documentation and examples for learning
- **Performance Optimized**: Benchmarked implementations with memory efficiency
- **Go Best Practices**: Idiomatic Go code with proper error handling

## 📚 Learning Path

### Phase 1: Fundamentals (📅 WIP)
- **Time & Space Complexity Analysis**
- **Big O Notation with practical examples**
- **Performance Benchmarking in Go**
- **Memory profiling and optimization**

### Phase 2: Data Structures (✅ Complete)

#### Linear Data Structures
- **Arrays & Slices**: Dynamic arrays with generic implementations
- **Stacks**: LIFO operations with multiple implementations
- **Queues**: FIFO operations including priority queues
- **Linked Lists**: Singly, doubly, and circular variants

## 🏗️ Project Structure

```
dsa-with-go/
├── 01-fundamentals/               # Complexity analysis and Go optimization
│   ├── complexity.go              # Time/space complexity examples
│   ├── benchmarking.go            # Performance measurement tools
│   ├── complexity_test.go         # Comprehensive tests
│   └── README.md                  # Fundamentals guide

```

## 🚀 Getting Started

### Prerequisites
- Go 1.21+ (uses generics for type safety)
- Basic understanding of Go syntax

### Installation

```bash
# Clone the repository
git clone <repository-url>
cd dsa-with-go

# Initialize Go module
go mod init dsa-with-go

# Run all tests to verify setup
go test ./...

# Run benchmarks
go test -bench=. -benchmem ./...
```


## 🧪 Testing & Benchmarking

### Running Tests

```bash
# Run all tests
go test ./...

# Run tests with verbose output
go test -v ./...

# Run specific package tests
go test -v ./01-fundamentals

# Run specific test
```

### Running Benchmarks

```bash
# Run all benchmarks
go test -bench=. -benchmem ./...

# Run specific benchmark

# Generate CPU profile
go tool pprof cpu.prof
```


## 📈 Performance Guidelines

### Benchmarking Best Practices
1. **Consistent Environment**: Run benchmarks on consistent hardware
2. **Multiple Runs**: Average results across multiple benchmark runs
3. **Memory Profiling**: Monitor allocations and GC pressure
4. **Real-world Data**: Test with realistic input sizes and patterns

### Optimization Strategies
1. **Algorithm Choice**: Select optimal algorithm for use case
2. **Data Structure Selection**: Match structure to access patterns
3. **Memory Management**: Minimize allocations and copying
4. **Concurrency**: Leverage Go's concurrency where beneficial

## 🤝 Contributing

### Code Standards
- Follow Go conventions and gofmt formatting
- Include comprehensive tests for new implementations
- Add benchmarks for performance-critical code
- Update documentation for new features

### Testing Requirements
- Unit tests with >95% coverage
- Benchmark tests for performance validation
- Edge case testing for robustness
- Example functions for documentation

## 📄 License

This project is created for educational purposes. Feel free to use, modify, and learn from the implementations.

## 🔗 Additional Resources
- [Go Documentation](https://golang.org/doc/)
- [Algorithm Visualization](https://visualgo.net/)
- [Big O Cheat Sheet](https://www.bigocheatsheet.com/)
- [Go Performance Tips](https://github.com/golang/go/wiki/Performance)


## 🛠 Tools & Resources
- Go 1.21+ for latest features
- Benchmarking tools for performance analysis
- Memory profiling for optimization
- Testing frameworks for validation

