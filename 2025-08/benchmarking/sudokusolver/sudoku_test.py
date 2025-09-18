#!/usr/bin/env python3
import time

from sudoku import example_board, solve_sudoku


def benchmark_sudoku_solver(rounds):
    # run test
    start_time = time.time()
    for _ in range(rounds):
        solve_sudoku(example_board)
    end_time = time.time()

    # calculate results
    total_time = end_time - start_time
    iterations_per_second = rounds / total_time
    op_time = (total_time / rounds) * 1000000000

    # display results
    print(
        f"benchmark_sudoku_solver\t{rounds:,}\t{total_time:.2f}s\t{iterations_per_second:,.2f} op/s\t\t{op_time:.2f} ns/op"
    )


if __name__ == "__main__":
    rounds = 10000000
    benchmark_sudoku_solver(rounds)
