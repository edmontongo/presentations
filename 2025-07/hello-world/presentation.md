---
marp: true
theme: default
class:
  - invert
paginate: true
---

<style>
section {
  padding-bottom: 20%;
}
</style>

# Hello, World!

## Your first Go program

Peter Preeper
2025-07-16
ppreeper@gmail.com

---

# Setting Up Go

- Mention [go.dev/dl](https://go.dev/dl) for downloads
- Briefly touch on `$GOPATH` and modules (`go mod init`)

---

# Your First Go Program: `main.go`

- Show basic structure

---

# Anatomy of `main.go`

- `package main` (executable program)
- `import "fmt"` (for formatted I/O)
- `func main() { ... }` (entry point)

---

# Basic Syntax: Variables

- Declaration: `var message string`
- Short declaration: `count := 10`
- Basic types: `int`, `string`, `bool`

---

# Printing to the Console

- `fmt.Println("Hello, Go!")`
- `fmt.Printf("Count: %d\n", count)`

---

# Basic Control Flow: `if` Statement

- `if condition { ... }`
- `if condition { ... } else { ... }`

---

# Basic Control Flow: `for` Loop

- `for i := 0; i < 5; i++ { ... }` (like C/Java)
- `for condition { ... }` (like `while`)
- `for { ... }` (infinite loop)

---

# Compiling and Running Your Program

- Compile: `go build` or `go build main.go`
- Run: `go run` or `go run main.go`
- Test: `go test`
- Benchmark: `go test -bench`

---

## Q&A

- Questions? Let’s discuss!

---

## About Me

- Name: Peter Preeper
- Contact: ppreeper@gmail.com
- I Work For: Thinksoft Inc.
- Job Title: Senior Implementation Specialist
