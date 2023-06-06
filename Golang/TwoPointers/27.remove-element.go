package main

// For array in-place operation , we have to consider if two pointers solution is applicable.
// Given (i, j) with initial value (0, 0)
// If nums[j] != target, we can move forward both i and j and assign nums[i] = nums[j] (to override the value when i !== j && nums[i] == target)
// Otherwise, we can only move j forward
// ![Demostration](https://camo.githubusercontent.com/3416a4d2775067bebeb7fe40955f4e9a59f6281c050bad8266853ac218b04685/68747470733a2f2f636f64652d7468696e6b696e672e63646e2e626365626f732e636f6d2f676966732f3937372e2545362539432538392545352542412538462545362539352542302545372542422538342545372539412538342545352542392542332545362539362542392e676966)

func removeElement(nums []int, val int) int {
	i, j := 0, 0
	length := len(nums)
	for j < length {
		if val != nums[j] {
			nums[i] = nums[j]
			i++
		}
		j++
	}
	return i
}
