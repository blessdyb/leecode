package main

func sortString(s string) string {
	data := make([]int, 26)
	for i := 0; i < len(s); i++ {
		data[s[i]]++
	}
	ret := make([]byte, len(s))
	for len(ret) < len(s) {
		lowest := 0
		for lowest < 26 {
			for lowest < 26 && data[lowest] == 0 {
				lowest++
			}
			ret = append(ret, byte(lowest))
			data[lowest]--
			lowest++
		}
		largest := 25
		for largest >= 0 {
			for largest >= 0 && data[largest] == 0 {
				largest--
			}
			ret = append(ret, byte(largest))
			data[largest]--
			largest--
		}
	}
	for i := 0; i < len(ret); i++ {
		ret[i] += 'a'
	}
	return string(ret)
}
