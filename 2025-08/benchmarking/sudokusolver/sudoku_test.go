package main

import "testing"

func TestSudokuSolver(t *testing.T) {
	if !solve_sudoku(example_board) {
		t.Errorf("Expected solved sudoku board, but got unsolved")
	}
	if !compareBoards(example_board, result_board) {
		t.Errorf("Expected solved sudoku board, but got %v", example_board)
	}
}

func BenchmarkCompareBoards(b *testing.B) {
	for b.Loop() {
		compareBoards(result_board, result_board)
	}
}

func BenchmarkSudokuSolverNoCheck(b *testing.B) {
	for b.Loop() {
		solve_sudoku(example_board)
	}
}

func BenchmarkSudokuSolver(b *testing.B) {
	for b.Loop() {
		solve_sudoku(example_board)
		if !compareBoards(example_board, result_board) {
			break
		}
	}
}

func BenchmarkSudokuSolverParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			solve_sudoku(example_board)
			// if !compareBoards(example_board, result_board) {
			// 	break
			// }
		}
	})
}
