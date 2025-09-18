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

# Benchmarking in Go

## How fast does it go?

Peter Preeper
2025-08-20
ppreeper@gmail.com

---

# Setting Up a Benchmark

- Need and appropriate test

---

# Your Go Program: `main.go`

- Show basic structure

---

# Layout of `main_test.go`

- `package main` (executable program)
- `import "testing"` (for testing library)
- `func TestSudokuSolver(t *testing.T) { ... }` (testing)
- `func BenchmarkSudokuSolver(b *testing.B) { ... }` (benchmark)
- `func BenchmarkSudokuSolverParallel(b *testing.B) { ... }` (benchmark)

---

# TestSudokuSolver

```go
func TestSudokuSolver(t *testing.T) {
	if !solve_sudoku(example_board) {
		t.Errorf("Expected solved sudoku board, but got unsolved")
	}
	if !compareBoards(example_board, result_board) {
		t.Errorf("Expected solved sudoku board, but got %v", example_board)
	}
}
```

---

# BenchmarkSudokuSolver

```go
func BenchmarkSudokuSolver(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solve_sudoku(example_board)
		if !compareBoards(example_board, result_board) {
			break
		}
	}
}
```

---

# BenchmarkSudokuSolverParallel

```go
func BenchmarkSudokuSolverParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			solve_sudoku(example_board)
			if !compareBoards(example_board, result_board) {
				break
			}
		}
	})
}
```

---

# Results

---

## Q&A

- Questions? Let’s discuss!

---

## About Me

- Name: Peter Preeper
- Contact: ppreeper@gmail.com
- I Work For: Thinksoft Inc.
- Job Title: Senior Implementation Specialist
