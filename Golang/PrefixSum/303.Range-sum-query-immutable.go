package main

type NumArray struct {
	preSum []int
	nums   []int
}

func Constructor(nums []int) NumArray {
	preSum := make([]int, len(nums))
	preSum[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		preSum[i] = nums[i] + preSum[i-1]
	}
	return NumArray{
		preSum: preSum,
		nums:   nums,
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.preSum[right] - this.preSum[left] + this.nums[left]
}
