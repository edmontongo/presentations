package main

import (
	"fmt"
)

var example_board = [][]int{
	{3, 9, 0, 0, 5, 0, 0, 0, 0},
	{0, 0, 0, 2, 0, 0, 0, 0, 5},
	{0, 0, 0, 7, 1, 9, 0, 8, 0},
	{0, 5, 0, 0, 6, 8, 0, 0, 0},
	{2, 0, 6, 0, 0, 3, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 4},
	{5, 0, 0, 0, 0, 0, 0, 0, 0},
	{6, 7, 0, 1, 0, 5, 0, 4, 0},
	{1, 0, 9, 0, 0, 0, 2, 0, 0},
}

var result_board = [][]int{
	{3, 9, 1, 8, 5, 6, 4, 2, 7},
	{8, 6, 7, 2, 3, 4, 9, 1, 5},
	{4, 2, 5, 7, 1, 9, 6, 8, 3},
	{7, 5, 4, 9, 6, 8, 1, 3, 2},
	{2, 1, 6, 4, 7, 3, 5, 9, 8},
	{9, 3, 8, 5, 2, 1, 7, 6, 4},
	{5, 4, 3, 6, 9, 2, 8, 7, 1},
	{6, 7, 2, 1, 8, 5, 3, 4, 9},
	{1, 8, 9, 3, 4, 7, 2, 5, 6},
}

func find_next_empty(puzzle [][]int) (int, int) {
	// finds the next row, col on the puzzle that's not filled yet --> rep with 0
	// return row, col tuple (or (-1,-1) if there is none)

	// keep in mind that we are using 0-8 for our indices
	for r := range 9 {
		for c := range 9 {
			if puzzle[r][c] == 0 {
				return r, c
			}
		}
	}
	// if no spaces in the puzzle are empty (0)
	return -1, -1
}

func is_valid(puzzle [][]int, guess int, row int, col int) bool {
	// figures out whether the guess at the row/col of the puzzle is a valid guess
	// returns True or False

	// for a guess to be valid, then we need to follow the sudoku rules
	// that number must not be repeated in the row, column, or 3x3 square that it appears in

	// let's start with the row
	row_vals := puzzle[row]
	for _, v := range row_vals {
		if v == guess {
			// if we've repeated, then our guess is not valid!
			return false
		}
	}

	// if slices.Contains(puzzle[row], guess) {
	// 	return false
	// }

	// for _, v := range puzzle[row] {
	// 	if v == guess {
	// 		return false
	// 	}
	// }

	// now the column
	col_vals := []int{}
	for r := range 9 {
		col_vals = append(col_vals, puzzle[r][col])
	}
	// Check if the guess is not in the current column
	for _, v := range col_vals {
		if v == guess {
			return false
		}
	}

	// for r := range 9 {
	// 	if puzzle[r][col] == guess {
	// 		return false
	// 	}
	// }

	// for r := range puzzle {
	// 	if puzzle[r][col] == guess {
	// 		return false
	// 	}
	// }

	// and then the square
	row_start := (row / 3) * 3
	col_start := (col / 3) * 3

	for r := row_start; r < row_start+3; r++ {
		for c := col_start; c < col_start+3; c++ {
			if puzzle[r][c] == guess {
				return false
			}
		}
	}

	return true
}

func solve_sudoku(puzzle [][]int) bool {
	// solve sudoku using backtracking!
	// our puzzle is a list of lists, where each inner list is a row in our sudoku puzzle
	// return whether a solution exists
	// mutates puzzle to be the solution (if solution exists)

	// step 1: choose somewhere on the puzzle to make a guess
	row, col := find_next_empty(puzzle)
	// step 1.1: if there's nowhere left, then we're done because we only allowed valid inputs
	// this is true if our find_next_empty function returns -1,-1
	if row == -1 {
		// Puzzle is solved
		return true
	}

	// step 2: if there is a place to put a number, then make a guess between 1 and 9
	for guess := 1; guess <= 9; guess++ {
		// step 3: check if this is a valid guess
		if is_valid(puzzle, guess, row, col) {
			// step 3.1: if this is a valid guess, then place it at that spot on the puzzle
			puzzle[row][col] = guess
			// step 4: then we recursively call our solver!
			if solve_sudoku(puzzle) {
				return true
			}

			// Backtrack
			// step 5: it not valid or if nothing gets returned true, then we need to backtrack and try a new number
			puzzle[row][col] = 0
		}
	}
	// step 6: if none of the numbers that we try work, then this puzzle is UNSOLVABLE!!
	return false
}

func main() {
	fmt.Println(solve_sudoku(example_board))
	for _, row := range example_board {
		fmt.Printf("%v\n", row)
	}
}

func compareBoards(board1, board2 [][]int) bool {
	for r := range 9 {
		for c := range 9 {
			if board1[r][c] != board2[r][c] {
				return false
			}
		}
	}
	return true
}
