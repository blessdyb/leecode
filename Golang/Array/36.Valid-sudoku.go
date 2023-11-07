package main

func isValidSudoku(board [][]byte) bool {
	isValid := func(cells []byte) bool {
		hashmap := map[byte]int{}
		for i := 0; i < len(cells); i++ {
			if cells[i] != '.' {
				hashmap[cells[i]]++
				if hashmap[cells[i]] > 1 {
					return false
				}
			}
		}
		return true
	}
	for i := 0; i < 9; i++ {
		if !isValid(board[i]) {
			return false
		}
		temp := []byte{}
		for j := 0; j < 9; j++ {
			temp = append(temp, board[j][i])
		}
		if !isValid(temp) {
			return false
		}
	}
	for i := 1; i < 9; i += 3 {
		for j := 1; j < 9; j += 3 {
			temp := []byte{}
			temp = append(temp, board[i][j])
			temp = append(temp, board[i-1][j-1])
			temp = append(temp, board[i+1][j+1])
			temp = append(temp, board[i+1][j-1])
			temp = append(temp, board[i-1][j+1])
			temp = append(temp, board[i-1][j])
			temp = append(temp, board[i][j-1])
			temp = append(temp, board[i+1][j])
			temp = append(temp, board[i][j+1])

			if !isValid(temp) {
				return false
			}
		}
	}
	return true
}
