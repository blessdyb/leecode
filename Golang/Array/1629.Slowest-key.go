package main

func slowestKey(releaseTimes []int, keysPressed string) byte {
	maxIndex := 0
	maxValue := releaseTimes[0]
	for i := 1; i < len(releaseTimes); i++ {
		if releaseTimes[i]-releaseTimes[i-1] > maxValue {
			maxIndex = i
			maxValue = releaseTimes[i] - releaseTimes[i-1]
		} else if releaseTimes[i]-releaseTimes[i-1] == maxValue && keysPressed[i] > keysPressed[maxIndex] {
			maxIndex = i
			maxValue = releaseTimes[i] - releaseTimes[i-1]
		}
	}
	return keysPressed[maxIndex]
}
