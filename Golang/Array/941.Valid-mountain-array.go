package main

func validMountainArray(arr []int) bool {
	if len(arr) < 3 || arr[1] <= 0 {
		return false
	}
	up := true
	down := false
	for i := 2; i < len(arr); i++ {
		if arr[i] == arr[i-1] {
			return false
		} else if arr[i] < arr[i-1] && up {
			up = false
			down = true
		} else if arr[i] > arr[i-1] && down {
			return false
		}
	}
	return !up && down
}
