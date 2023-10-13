package main

func judgeCircle(moves string) bool {
	i, j := 0, 0
	for k := 0; k < len(moves); k++ {
		if moves[k] == 'U' {
			j++
		} else if moves[k] == 'D' {
			j--
		} else if moves[k] == 'L' {
			i--
		} else {
			i++
		}
	}
	return i == 0 && j == 0
}
