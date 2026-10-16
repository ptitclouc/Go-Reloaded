package main

func Capitalize(s string) string {
	result := ""
	newWord := true
	for _, capi := range s {
		if !((capi >= 'a' && capi <= 'z') || (capi >= 'A' && capi <= 'Z') || (capi >= '0' && capi <= '9')) {
			newWord = true
			result += string(capi)
			continue
		}
		if newWord {
			if capi >= 'a' && capi <= 'z' {
				capi -= 32
			}
			newWord = false
		} else {
			if capi >= 'A' && capi <= 'Z' {
				capi += 32
			}
		}
		result += string(capi)
	}
	return result
}
