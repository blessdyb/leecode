package main

func strongPasswordCheckerII(password string) bool {
	if len(password) < 8 {
		return false
	}
	hasLowercase, hasUppercase, hasDigit, hasSpecialChars, noRepeat := false, false, false, false, true
	for i := 0; i < len(password); i++ {
		if i > 0 {
			if password[i] == password[i-1] {
				noRepeat = false
			}
		}
		if password[i] >= 'a' && password[i] <= 'z' {
			hasLowercase = true
		} else if password[i] >= 'A' && password[i] <= 'Z' {
			hasUppercase = true
		} else if password[i] >= '0' && password[i] <= '9' {
			hasDigit = true
		} else {
			hasSpecialChars = true
		}
	}
	return hasLowercase && hasUppercase && hasDigit && hasSpecialChars && noRepeat
}
