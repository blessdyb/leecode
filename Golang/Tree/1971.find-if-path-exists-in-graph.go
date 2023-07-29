package main

func validPathDFS(n int, edges [][]int, source int, destination int) bool {
	hashmap := map[int][]int{}
	for i := 0; i < n; i++ {
		hashmap[i] = []int{}
	}
	for _, e := range edges {
		hashmap[e[0]] = append(hashmap[e[0]], e[1])
		hashmap[e[1]] = append(hashmap[e[1]], e[0])
	}
	visited := map[int]bool{}
	var helper func(src int) bool
	helper = func(src int) bool {
		if src == destination {
			return true
		}
		if !visited[src] {
			visited[src] = true
			for _, dst := range hashmap[src] {
				if helper(dst) {
					return true
				}
			}
		}
		return false
	}
	return helper(source)
}

func validPathBFS(n int, edges [][]int, source int, destination int) bool {
	hashmap := map[int][]int{}
	for i := 0; i < n; i++ {
		hashmap[i] = []int{}
	}
	for _, e := range edges {
		hashmap[e[0]] = append(hashmap[e[0]], e[1])
		hashmap[e[1]] = append(hashmap[e[1]], e[0])
	}
	visited := map[int]bool{}
	stack := []int{source}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node == destination {
			return true
		}
		if !visited[node] {
			visited[node] = true
			stack = append(stack, hashmap[node]...)
		}
	}
	return false
}
