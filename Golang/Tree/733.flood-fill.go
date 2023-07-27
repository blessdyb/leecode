package main

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	sourceColor := image[sr][sc]
	row, col := len(image), len(image[0])
	var helper func(r, c int)
	helper = func(r, c int) {
		if image[r][c] == sourceColor {
			image[r][c] = color
			if r >= 1 {
				helper(r-1, c)
			}
			if c >= 1 {
				helper(r, c-1)
			}
			if r+1 < row {
				helper(r+1, c)
			}
			if c+1 < col {
				helper(r, c+1)
			}
		}
	}
	if sourceColor != color {
		helper(sr, sc)
	}
	return image
}
