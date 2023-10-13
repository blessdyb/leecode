package main

func licenseKeyFormatting(s string, k int) string {
	ret := ""
	temp := ""
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != '-' {
			c := s[i]
			if c >= 'a' && c <= 'z' {
				c = c - 'a' + 'A'
			}
			temp = string(c) + temp
		}
		if len(temp) == k || (i == 0 && len(temp) > 0) {
			if ret == "" {
				ret = temp
			} else {
				ret = temp + "-" + ret
			}
			temp = ""
		}
	}
	return ret
}
