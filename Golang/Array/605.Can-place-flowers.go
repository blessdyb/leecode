package main

func canPlaceFlowers(flowerbed []int, n int) bool {
	for i := 0; i < len(flowerbed) && n > 0; i++ {
		if flowerbed[i] == 1 {
			i++
		} else if flowerbed[i] == 0 {
			if (i == len(flowerbed)-1 || flowerbed[i+1] == 0) && (i == 0 || flowerbed[i-1] == 0) {
				n--
				flowerbed[i] = 1
			}
		}
	}
	return n == 0
}
