package main

type Heap struct {
	Data     []int
	Capacity int
}

func NewHeap(capacity int) *Heap {
	return &Heap{
		Data:     []int{},
		Capacity: capacity,
	}
}

func (this *Heap) Push(value int) {
	if len(this.Data) < this.Capacity {
		this.Data = append(this.Data, value)
		index := len(this.Data) - 1
		parentIndex := (index - 1) / 2
		for index >= 0 && this.Data[parentIndex] > this.Data[index] {
			this.Data[parentIndex], this.Data[index] = this.Data[index], this.Data[parentIndex]
			index = parentIndex
			parentIndex = (index - 1) / 2
		}
	} else if this.Data[0] < value {
		this.Data[0] = value
		index := 0
		left := 2*index + 1
		for left < this.Capacity {
			child := left
			if child+1 < this.Capacity && this.Data[child+1] < this.Data[child] {
				child++
			}
			if this.Data[child] < this.Data[index] {
				this.Data[child], this.Data[index] = this.Data[index], this.Data[child]
				index = child
				left = 2*index + 1
			} else {
				break
			}
		}
	}
}

func kthLargestValue(matrix [][]int, k int) int {
	dp := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		dp[i] = make([]int, len(matrix[0]))
	}
	heap := NewHeap(k)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if i == 0 && j == 0 {
				dp[0][0] = matrix[0][0]
			} else if i == 0 {
				dp[i][j] = dp[i][j-1] ^ matrix[i][j]
			} else if j == 0 {
				dp[i][j] = dp[i-1][j] ^ matrix[i][j]
			} else {
				dp[i][j] = dp[i-1][j-1] ^ dp[i-1][j] ^ dp[i][j-1] ^ matrix[i][j]
			}
			heap.Push(dp[i][j])
		}
	}
	return heap.Data[0]
}
