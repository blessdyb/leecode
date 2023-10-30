package main

func squareIsWhite(coordinates string) bool {
	x := (coordinates[0] - 'a') % 2
	y := (coordinates[1] - '1') % 2
	return !((x == 0 && y == 0) || (x == 1 && y == 1))
}
