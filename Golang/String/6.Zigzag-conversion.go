package main

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	matrix := [][]byte{}
	for i := 0; i < numRows; i++ {
		matrix = append(matrix, []byte{})
	}
	for i, j, flag := 0, 0, true; i < len(s); i++ {
		matrix[j] = append(matrix[j], s[i])
		if flag {
			j++
		} else {
			j--
		}
		if j == numRows {
			j = numRows - 2
			flag = false
		} else if j == -1 {
			j = 1
			flag = true
		}
	}
	ret := []byte{}
	for i := 0; i < numRows; i++ {
		ret = append(ret, matrix[i]...)
	}
	return string(ret)
}
