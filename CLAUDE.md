# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a Go-based algorithms and data structures repository containing implementations of common algorithms organized by category. The codebase includes:

- **Algorithms**: Organized by type (sorting, graph, dynamic programming, etc.) in `/algorithms/`
- **Data Structures**: Core data structure definitions in `/data_structures/`
- **Engineering**: Specialized implementations like Bloom filters
- **Jupyter Notebooks**: Problem solutions organized by algorithm category

## Development Commands

**Build and Format:**
```bash
go build ./...          # Build all packages
go fmt ./...           # Format all Go code
```

**Testing:**
```bash
go test ./...          # Run all tests (currently no test files exist)
```

**Module Management:**
```bash
go mod tidy            # Clean up module dependencies
```

## Code Structure

**Main Packages:**
- `algorithms-go/algorithms` - Algorithm implementations
- `algorithms-go/data_structures` - Core data structures
- `algorithms-go/engineering` - Specialized implementations

**Key Data Structures:**
- `ListNode` - Singly linked list node (`data_structures/list_node.go:7`)
- `TreeNode` - Binary tree node (`data_structures/tree_node.go`)
- `GraphNode` - Graph node (`data_structures/graph_node.go`)

**Algorithm Categories:**
- Array manipulation, binary search, bit manipulation
- Dynamic programming, graph algorithms, greedy algorithms  
- Hash-based algorithms, math problems, prefix sums
- Recursion, sliding window, sorting, stack operations
- String algorithms, tree algorithms, two pointers

## Code Style

- Go 1.21+ syntax and conventions
- Chinese and English comments for algorithm descriptions
- Package-level organization by algorithm type
- Data structures defined separately from algorithms
- No existing test suite (opportunity to add tests)

## Common Development Tasks

1. **Adding new algorithms**: Place in appropriate category file in `/algorithms/`
2. **Extending data structures**: Modify files in `/data_structures/`
3. **Creating tests**: Add `_test.go` files alongside implementations
4. **Formatting**: Use `go fmt` before committing changes