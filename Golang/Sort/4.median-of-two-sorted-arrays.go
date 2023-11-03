package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	p, q, m, n := 0, 0, len(nums1), len(nums2)
	min := func() float64 {
		ret := 0
		if p < m && q < m {
			if nums1[p] < nums2[q] {
				ret = nums1[p]
				p++
			} else {
				ret = nums2[q]
				q++
			}
		} else if p < m {
			ret = nums1[p]
			p++
		} else {
			ret = nums2[q]
			q++
		}
		return float64(ret)
	}
	if (m+n)%2 == 0 {
		for i := 0; i < (m+n)/2-1; i++ {
			min()
		}
		return (min() + min()) / 2
	}
	for i := 0; i < (m+n)/2; i++ {
		min()
	}
	return min()
}
