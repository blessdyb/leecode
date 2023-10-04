package main

func numSpecial(mat [][]int) int {
	m, n := len(mat)-1, len(mat[0])-1
	ret := 0
	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			if mat[i][j] == 1 {
				temp := 0
				for l := 0; l <= n; l++ {
					temp += mat[i][l]
					if temp > 1 {
						break
					}
				}
				if temp == 1 {
					temp = 0
					for l := 0; l <= m; l++ {
						temp += mat[l][j]
						if temp > 1 {
							break
						}
					}
					if temp == 1 {
						ret++
					}
				}
			}
		}
	}
	return ret
}
