package main

func getRow(rowIndex int) []int {
	ret := make([]int, rowIndex+1)
	ret[0] = 1
	ret[rowIndex] = 1
	lastRow := []int{1, 1}
	for i := 2; i <= rowIndex; i++ {
		temp := make([]int, i+1)
		temp[0] = 1
		temp[i] = 1
		for j := 1; j < i; j++ {
			temp[j] = lastRow[j-1] + lastRow[j]
		}
		lastRow = temp
	}
	if rowIndex >= 2 {
		ret = lastRow
	}
	return ret
}
