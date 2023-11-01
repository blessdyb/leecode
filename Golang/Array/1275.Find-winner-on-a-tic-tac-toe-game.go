package main

func tictactoe(moves [][]int) string {
	board := [3][3]byte{}
	for i := 0; i < 3; i++ {
		board[i] = [3]byte{}
	}
	for i := 0; i < len(moves); i++ {
		if i%2 == 0 {
			board[moves[i][0]][moves[i][1]] = 'X'
		} else {
			board[moves[i][0]][moves[i][1]] = 'O'
		}
	}
	getResult := func(move byte) string {
		if move == 'X' {
			return "A"
		} else {
			return "B"
		}
	}
	for i := 0; i < 3; i++ {
		if board[i][0] == board[i][1] && board[i][0] == board[i][2] && board[i][0] != 0 {
			return getResult(board[i][0])
		}
		if board[0][i] == board[1][i] && board[0][i] == board[2][i] && board[0][i] != 0 {
			return getResult(board[0][i])
		}
	}
	if board[0][0] == board[1][1] && board[0][0] == board[2][2] && board[0][0] != 0 {
		return getResult(board[0][0])
	}
	if board[2][0] == board[1][1] && board[2][0] == board[0][2] && board[1][1] != 0 {
		return getResult(board[2][0])
	}
	if len(moves) < 9 {
		return "Pending"
	}
	return "Draw"
}
