package main

func hardestWorker(n int, logs [][]int) int {
	ret := logs[0][0]
	max := logs[0][1]
	for i := 1; i < len(logs); i++ {
		if logs[i][1]-logs[i-1][1] > max {
			ret = logs[i][0]
			max = logs[i][1] - logs[i-1][1]
		} else if logs[i][1]-logs[i-1][1] == max && logs[i][0] < ret {
			ret = logs[i][0]
		}
	}
	return ret
}
