package main

func oddCells(m int, n int, indices [][]int) int {
	rows := make([]int, m)
	cols := make([]int, n)
	for _, index := range indices {
		rows[index[0]]++
		cols[index[1]]++
	}
	oddRows, oddCols := 0, 0
	for _, row := range rows {
		if row%2 == 1 {
			oddRows++
		}
	}
	for _, col := range cols {
		if col%2 == 1 {
			oddCols++
		}
	}
	return m*oddCols + n*oddRows - 2*oddCols*oddRows
}
