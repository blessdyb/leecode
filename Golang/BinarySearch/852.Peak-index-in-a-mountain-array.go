package main

func peakIndexInMountainArray(arr []int) int {
	ret := -1
	for i, j := 0, len(arr)-1; i <= j; {
		mid := (j-i)/2 + i
		if arr[mid] > arr[mid-1] && arr[mid] > arr[mid+1] {
			ret = mid
			break
		} else if arr[mid] > arr[mid-1] && arr[mid] < arr[mid+1] {
			i = mid
		} else {
			j = mid
		}
	}
	return ret
}
