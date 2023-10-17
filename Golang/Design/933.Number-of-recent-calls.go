package main

type RecentCounter struct {
	data []int
}

func Constructor() RecentCounter {
	return RecentCounter{
		data: []int{},
	}
}

func (this *RecentCounter) Ping(t int) int {
	this.data = append(this.data, t)
	count := 0
	for i := len(this.data) - 1; i >= 0; i-- {
		if this.data[i] <= t && this.data[i] >= t-3000 {
			count++
		} else {
			break
		}
	}
	return count
}
