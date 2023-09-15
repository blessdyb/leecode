package main

func merge(nums1 []int, m int, nums2 []int, n int)  {
	if n != 0 {
		for i, j, k := n - 1, m - 1, m + n - 1; i >= 0 || j >= 0 {
			if i >=0 && j >= 0 && nums1[j] > nums2[j] {
				nums1[k] = nums1[j]
				j--
			} else if i >= 0 {
				nums1[k] = nums2[i]
				i--
			} else {
				j--
			}
			k--
		}
	}
}