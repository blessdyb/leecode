package main

func defangIPaddr(address string) string {
	ret := []byte{}
	for i := 0; i < len(address); i++ {
		if address[i] == '.' {
			ret = append(ret, []byte{'[', '.', ']'}...)
		} else {
			ret = append(ret, address[i])
		}
	}
	return string(ret)
}
