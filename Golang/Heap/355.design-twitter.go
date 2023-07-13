package main

type Heap struct {
	data     [][2]int
	capacity int
}

func NewHeap(c int) Heap {
	return Heap{
		data:     [][2]int{},
		capacity: c,
	}
}

func (this *Heap) down() {
	parentIndex := 0
	left := parentIndex*2 + 1
	for left < this.Len() {
		child := left
		if child+1 < this.Len() && this.data[child+1][1] < this.data[child][1] {
			child = child + 1
		}
		if this.data[child][1] < this.data[parentIndex][1] {
			this.data[parentIndex], this.data[child] = this.data[child], this.data[parentIndex]
			parentIndex = child
			left = parentIndex*2 + 1
		} else {
			break
		}
	}
}

func (this *Heap) Pop() int {
	var ret int
	if this.Len() > 0 {
		ret = this.data[0][0]
		this.data[0] = this.data[this.Len()-1]
		this.data = this.data[:this.Len()-1]
		this.down()
		return ret
	}
	return ret
}

func (this *Heap) Len() int {
	return len(this.data)
}

func (this *Heap) Add(tweet [2]int) {
	if this.Len() >= this.capacity {
		if tweet[1] < this.data[0][1] {
			return
		} else {
			this.data[0] = tweet
			this.down()
		}
	} else {
		this.data = append(this.data, tweet)
		index := this.Len() - 1
		parentIndex := (index - 1) / 2
		for index >= 0 && this.data[parentIndex][1] > this.data[index][1] {
			this.data[parentIndex], this.data[index] = this.data[index], this.data[parentIndex]
			index = parentIndex
			parentIndex = (index - 1) / 2
		}
	}
}

type Twitter struct {
	tweets   map[int][][2]int
	followee map[int]map[int]bool
	id       int
}

func Constructor() Twitter {
	return Twitter{
		tweets:   map[int][][2]int{},
		followee: map[int]map[int]bool{},
		id:       0,
	}
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	this.tweets[userId] = append(this.tweets[userId], [2]int{tweetId, this.id})
	this.id++
}

func (this *Twitter) GetNewsFeed(userId int) []int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	heap := NewHeap(10)
	for _, tweet := range this.tweets[userId][max(0, len(this.tweets[userId])-10):] {
		heap.Add(tweet)
	}
	for followee, _ := range this.followee[userId] {
		followeeTweets := this.tweets[followee]
		for i := len(followeeTweets) - 1; i >= 0; i-- {
			heap.Add(followeeTweets[i])
		}
	}
	ret := make([]int, heap.Len())
	for i := heap.Len() - 1; i >= 0; i-- {
		ret[i] = heap.Pop()
	}
	return ret
}

func (this *Twitter) Follow(followerId int, followeeId int) {
	if followerId != followeeId {
		if _, ok := this.followee[followerId]; !ok {
			this.followee[followerId] = map[int]bool{}
		}
		this.followee[followerId][followeeId] = true
	}
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	if followerId != followeeId {
		delete(this.followee[followerId], followeeId)
	}
}
