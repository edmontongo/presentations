test:
	go test -v

benchmark:
	go test -bench=.

benchmem:
	go test -bench=. -benchmem

benchcpu:
	go test -bench=. -benchmem -cpu 1,2,4,8

benchcount:
	go test -bench=. -benchmem -count=4

benchpy:
	python3 sudoku_test.py