/** Given any integer, it's a sorted digits 1 ~ n and we want to find sqrt(x) which is a perfect case for the binary search
 */
func mySqrt(x int) int {
	i, j := 0, x
	ret := 0
	for i <= j {
		h := i + (j - i) / 2
		h2 = h * h
		if h2 === x {
			ret = h
			break
		} else if h2 < x {
			i = h + 1
		} else {
			j = h - 1
		}
	}
	return ret
}