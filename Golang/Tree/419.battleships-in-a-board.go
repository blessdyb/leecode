package main

func countBattleships(board [][]byte) int {
	m, n := len(board), len(board[0])
	inBound := func(r, c int) bool {
		return r >= 0 && r < m && c >= 0 && c < n
	}
	var dfs func(r, c int)
	dfs = func(r, c int) {
		if !inBound(r, c) || board[r][c] == '.' {
			return
		}
		board[r][c] = '.'
		dfs(r+1, c)
		dfs(r-1, c)
		dfs(r, c+1)
		dfs(r, c-1)
	}
	ret := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'X' {
				ret++
				dfs(i, j)
			}
		}
	}
	return ret
}

func countBattleshipsOneRun(board [][]byte) int {
	ret := 0
	m, n := len(board), len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// No ship on the top right direction
			if board[i][j] == 'X' && (i == 0 || board[i-1][j] == '.') && (j == 0 || board[i][j-1] == '.') {
				ret++
			}
		}
	}
	return ret
}
