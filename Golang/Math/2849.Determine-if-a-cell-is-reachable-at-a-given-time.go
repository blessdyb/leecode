package main

func isReachableAtTime(sx int, sy int, fx int, fy int, t int) bool {
	abs := func(a int) int {
		if a >= 0 {
			return a
		}
		return -a
	}
	dx, dy := abs(fx-sx), abs(fy-sy)
	if dx == 0 && dy == 0 && t == 1 {
		return false
	}
	if dx == 0 || dy >= dx {
		return dy <= t
	}
	if dy == 0 || dx >= dy {
		return dx <= t
	}
	return false
}
