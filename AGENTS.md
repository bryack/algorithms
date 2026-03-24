# Project Overview

This is a **personal algorithm learning project** written in Go. It serves as a structured study repository for mastering data structures and algorithms, specifically designed for preparing for technical interviews at big tech companies.

The project follows a comprehensive 6-month study plan (documented in `Plan. Introduction to Algorithms.md`) that covers mathematical foundations, array/string patterns, linear data structures, trees/graphs, and dynamic programming.

**Note**: Project documentation and comments are primarily in Russian.

## Technology Stack

- **Language**: Go 1.25.8
- **Module Path**: `github.com/bryack/algorithms`
- **Build System**: Standard Go toolchain (`go build`, `go test`, `go run`)

## Project Structure

```
.
├── go.mod                          # Go module definition
├── README.md                       # Empty (to be populated)
├── Plan. Introduction to Algorithms.md  # Main study plan (Russian)
├── AGENTS.md                       # This file
├── .gitignore                      # Git ignore rules
└── .serena/                        # Serena AI assistant configuration
    ├── project.yml                 # Project settings for Serena
    └── memories/                   # Serena memory storage (empty)
```

## Code Organization

The project is currently in its initial setup phase. Source code is expected to be organized by topic according to the study plan:

1. **Arrays and Strings** - Hash maps, two pointers, sliding window patterns
2. **Linear Structures** - Stacks, queues, linked lists
3. **Recursion & Divide and Conquer** - Binary search, merge/quick sort
4. **Trees & Graphs** - BST, BFS, DFS traversals
5. **Dynamic Programming** - Memoization, tabulation patterns

## Build and Test Commands

```bash
# Build the project
go build ./...

# Run tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run a specific file
go run <filename>.go

# Format code
go fmt ./...

# Vet code for issues
go vet ./...

# Download dependencies
go mod download

# Tidy module dependencies
go mod tidy
```

## Code Style Guidelines

- Follow standard Go conventions (`gofmt`, `golint`)
- Use meaningful variable names in English
- Include comments explaining algorithm complexity: `// Time: O(n), Space: O(1)`
- Prefer iterative solutions over recursive for deep recursion cases (avoid stack overflow)
- Handle edge cases explicitly (empty inputs, single elements)

## Go-Specific Considerations for Algorithms

Based on the study plan, keep these Go specifics in mind:

1. **Strings are bytes, not runes**: For Unicode problems, convert to `[]rune` instead of `[]byte`
2. **Slices are passed by value, but share underlying array**: Be careful with side effects when modifying input
3. **Map iteration order is randomized**: Use a slice of keys if order is required
4. **No built-in Set**: Implement using `map[int]struct{}` (empty struct uses zero memory)
5. **Nil pointer safety**: Always check for `nil` when working with linked lists and trees

## Testing Strategy

- Write unit tests for each algorithm implementation
- Include test cases for:
  - Normal cases
  - Edge cases (empty input, single element)
  - Boundary conditions
  - Performance stress tests where applicable
- Test files should follow Go convention: `<filename>_test.go`

## External Resources

- **LeetCode**: Primary platform for practice (filter: "Top Interview 150")
- **Go Playground**: For quick experiments without IDE
- **Khan Academy / 3Blue1Brown**: Mathematical foundations

## Development Workflow

1. Pick a topic from the study plan
2. Implement the algorithm/data structure
3. Write comprehensive tests
4. Document complexity analysis
5. Commit with descriptive messages

## Serena Configuration

The project is configured for use with Serena AI assistant:
- Project name: `algorithms`
- Language backend: LSP (Language Server Protocol)
- Encoding: UTF-8
- Read-only mode: Disabled
