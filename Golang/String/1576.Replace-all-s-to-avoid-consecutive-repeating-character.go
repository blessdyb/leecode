package main

func modifyString(s string) string {
	chars := []byte(s)
	for i := 0; i < len(chars); i++ {
		if chars[i] == '?' {
			if i == 0 {
				var temp byte = 'a'
				for i+1 < len(chars) && chars[i+1] == temp {
					temp = byte(temp + 1)
				}
				chars[i] = temp
			} else if i == len(chars)-1 {
				var temp byte = 'a'
				for i-1 < len(chars) && chars[i-1] == temp {
					temp = byte(temp + 1)
				}
				chars[i] = temp
			} else {
				var temp byte = 'a'
				for chars[i+1] == temp || chars[i-1] == temp {
					temp = byte(temp + 1)
				}
				chars[i] = temp
			}
		}
	}
	return string(chars)
}
