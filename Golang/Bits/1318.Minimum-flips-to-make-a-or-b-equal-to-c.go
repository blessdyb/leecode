package main

func minFlips(a int, b int, c int) int {
	getDigits := func(num int) []int {
		digits := []int{}
		for num > 0 {
			digits = append(digits, num&1)
			num >>= 1
		}
		return digits
	}
	count := 0
	aBytes, bBytes, cBytes := getDigits(a), getDigits(b), getDigits(c)
	for i, j, k := 0, 0, 0; i < len(aBytes) || j < len(bBytes) || k < len(cBytes); i, j, k = i+1, j+1, k+1 {
		aDigit, bDigit, cDigit := 0, 0, 0
		if i < len(aBytes) {
			aDigit = aBytes[i]
		}
		if j < len(bBytes) {
			bDigit = bBytes[j]
		}
		if k < len(cBytes) {
			cDigit = cBytes[k]
		}
		if (aDigit | bDigit) != cDigit {
			if cDigit == 1 {
				count += 1
			} else {
				if aDigit == bDigit {
					count += 2
				} else {
					count += 1
				}
			}
		}
	}
	return count
}
