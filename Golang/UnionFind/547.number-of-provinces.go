package main

type UnionFind struct {
	Parent []int
	Rank   []int
}

func NewUnionFind(size int) *UnionFind {
	uf := &UnionFind{
		Parent: make([]int, size),
		Rank:   make([]int, size),
	}
	for i := 0; i < size; i++ {
		uf.Parent[i] = i
		uf.Rank[i] = 1
	}
	return uf
}

func (this *UnionFind) Find(x int) int {
	if this.Parent[x] != x {
		// Path compression
		this.Parent[x] = this.Find(this.Parent[x])
	}
	return this.Parent[x]
}

func (this *UnionFind) Union(x, y int) {
	rootX := this.Find(x)
	rootY := this.Find(y)
	if rootX != rootY {
		if this.Rank[rootX] > this.Rank[rootY] {
			this.Parent[rootY] = rootX
		} else if this.Rank[rootX] < this.Rank[rootY] {
			this.Parent[rootX] = rootY
		} else {
			this.Parent[rootY] = rootX
			this.Rank[rootX]++
		}
	}
}

func findCircleNumUF(isConnected [][]int) int {
	size := len(isConnected)
	ret := size
	uf := NewUnionFind(size)
	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			if isConnected[i][j] == 1 && uf.Find(i) != uf.Find(j) {
				uf.Union(i, j)
				ret--
			}
		}
	}
	return ret
}
