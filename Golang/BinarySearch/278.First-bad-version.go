package main

func isBadVersion(version int) bool

func firstBadVersion(n int) int {
	ret := 1
	for i, j := 1, n; i <= j; {
		mid := (j-i)/2 + i
		if !isBadVersion(mid) {
			i = mid + 1
		} else if isBadVersion((mid - 1)) {
			j = mid - 1
		} else {
			ret = mid
			break
		}
	}
	return ret
}
