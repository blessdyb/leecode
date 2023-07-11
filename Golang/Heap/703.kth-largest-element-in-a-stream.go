package main

type KthLargest struct {
	nums []int
	k    int
}

func Constructor(k int, nums []int) KthLargest {
	kthLargest := &KthLargest{
		nums: []int{},
		k:    k,
	}
	for _, num := range nums {
		kthLargest.Add(num)
	}
	return *kthLargest
}

func (this *KthLargest) Add(val int) int {
	if len(this.nums) < this.k {
		this.nums = append(this.nums, val)
		index := len(this.nums) - 1
		parentIndex := (index - 1) / 2
		for parentIndex >= 0 && this.nums[parentIndex] > this.nums[index] {
			this.nums[parentIndex], this.nums[index] = this.nums[index], this.nums[parentIndex]
			index = parentIndex
			parentIndex = (index - 1) / 2
		}
	} else if this.nums[0] < val {
		this.nums[0] = val
		parentIndex := 0
		left := parentIndex*2 + 1
		for left < this.k {
			child := left
			if child+1 < this.k && this.nums[child+1] < this.nums[child] {
				child = child + 1
			}
			if this.nums[child] < this.nums[parentIndex] {
				this.nums[parentIndex], this.nums[child] = this.nums[child], this.nums[parentIndex]
				parentIndex = child
				left = parentIndex*2 + 1
			} else {
				break
			}
		}
	}
	return this.nums[0]
}
