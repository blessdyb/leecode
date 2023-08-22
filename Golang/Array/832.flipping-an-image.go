package main

func flipAndInvertImage(image [][]int) [][]int {
	for i := 0; i < len(image); i++ {
		for j, k := 0, len(image)-1; j <= k; {
			temp := image[i][j] ^ 1
			image[i][j] = image[i][k] ^ 1
			image[i][k] = temp
			j++
			k--
		}
	}
	return image
}
