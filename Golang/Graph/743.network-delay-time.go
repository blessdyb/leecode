package main

func networkDelayTime(times [][]int, n int, k int) int {
	timeSpent := map[int]int{}
	visited := map[int]bool{}
	nodes := map[int]map[int]int{}
	for _, time := range times {
		if nodes[time[0]] == nil {
			nodes[time[0]] = map[int]int{}
		}
		nodes[time[0]][time[1]] = time[2]
	}
	var help func(dst, ts int)
	help = func(dst, ts int) {
		visited[dst] = true
		if _, ok := timeSpent[dst]; !ok {
			timeSpent[dst] = ts
		} else if timeSpent[dst] > ts {
			timeSpent[dst] = ts
		} else {
			return
		}
		for k, w := range nodes[dst] {
			help(k, timeSpent[dst]+w)
		}
	}
	help(k, 0)
	ret := -1
	for i := 1; i <= n; i++ {
		if !visited[i] {
			return ret
		}
	}
	for _, v := range timeSpent {
		if ret < v {
			ret = v
		}
	}
	return ret
}
