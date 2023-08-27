package main

import "fmt"

func numRookCaptures(board [][]byte) int {
	var rook []int
	m, n := len(board), len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'R' {
				rook = []int{i, j}
			}
		}
	}
	ret := 0
	k, l, p, q := rook[1]-1, rook[1]+1, rook[0]-1, rook[0]+1
	for k >= 0 || l < n || p >= 0 || q < n {
		if k >= 0 {
			if board[rook[0]][k] == 'p' {
				ret++
				k = -1
			} else if board[rook[0]][k] == '.' {
				k--
			} else if board[rook[0]][k] == 'B' {
				k = -1
			}
		}
		if l < n {
			fmt.Println(rook[1], l)
			if board[rook[0]][l] == 'p' {
				ret++
				l = n
			} else if board[rook[0]][l] == '.' {
				l++
			} else if board[rook[0]][l] == 'B' {
				l = n
			}
		}
		if p >= 0 {
			if board[p][rook[1]] == 'p' {
				ret++
				p = -1
			} else if board[p][rook[1]] == '.' {
				p--
			} else if board[p][rook[1]] == 'B' {
				p = -1
			}
		}
		if q < m {
			if board[q][rook[1]] == 'p' {
				ret++
				q = m
			} else if board[q][rook[1]] == '.' {
				q++
			} else if board[q][rook[1]] == 'B' {
				q = m
			}
		}
	}
	return ret
}
