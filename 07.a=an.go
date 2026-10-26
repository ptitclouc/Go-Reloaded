package main

func atooan(mots []string) []string {
	var result []string
	for _, m := range mots {
		c := string([]rune(m)[:1])
		if m == "a" && c == "a" || c == "e" || c == "i" || c == "o" || c == "u" || c == "y"{
			last := len(result) - 1
			result[last] += " an"
			continue
		}
		result = append(result, m)
	} 
	return result
}
