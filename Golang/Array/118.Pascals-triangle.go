package main

func generate(numRows int) [][]int {
	ret := [][]int{[]int{1}}
	for i := 2; i <= numRows; i++ {
		row := make([]int, i)
		row[0] = 1
		row[i-1] = 1
		lastRow := ret[len(ret)-1]
		for j := 1; j < i-1; j++ {
			row[j] = lastRow[j-1] + lastRow[j]
		}
		ret = append(ret, row)
	}
	return ret
}
