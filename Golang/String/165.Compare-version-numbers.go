package main

func compareVersion(version1, version2 string) int {
	getNumber := func(digits []int) int {
		ret := 0
		base := 1
		for i := len(digits) - 1; i >= 0; i-- {
			ret = ret + digits[i]*base
			base *= 10
		}
		return ret
	}
	getVersion := func(version string) []int {
		ret := []int{}
		temp := []int{}
		for i := 0; i < len(version); i++ {
			if version[i] == '.' {
				ret = append(ret, getNumber(temp))
				temp = []int{}
			} else {
				temp = append(temp, int(version[i]-'0'))
			}
		}
		ret = append(ret, getNumber(temp))
		return ret
	}
	v1, v2 := getVersion(version1), getVersion(version2)
	for i, j := 0, 0; i < len(v1) || j < len(v2); {
		if i < len(v1) && j < len(v2) {
			if v1[i] < v2[j] {
				return -1
			} else if v1[i] > v2[j] {
				return 1
			}
			i++
			j++
		} else if i < len(v1) {
			if v1[i] > 0 {
				return 1
			}
			i++
		} else {
			if v2[j] > 0 {
				return -1
			}
			j++
		}
	}
	return 0
}
