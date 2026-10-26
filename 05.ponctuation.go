package main

func ponctuation(mots []string) []string {
	var result []string
	for _, m := range mots {
		c := string([]rune(m)[:1])
		if (c == "." || c == "," || c == "!" || c == "?" || c == ":" || c == ";") && len(result) > 0 {
			result[len(result)-1] += m
			continue
		}
		result = append(result, m)
	}
	return result
}
