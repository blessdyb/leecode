package main

import "math"

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	ret := 0
	for i := 0; i < len(arr1); i++ {
		flag := true
		for j := 0; j < len(arr2); j++ {
			if int(math.Abs(float64(arr1[i]-arr2[j]))) <= d {
				flag = false
				break
			}
		}
		if flag {
			ret++
		}
	}
	return ret
}
