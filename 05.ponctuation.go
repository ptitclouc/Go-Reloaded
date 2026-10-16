package main

func ponctuation(mots []string) []string{
	var result []string
	for i, m:= range mots {
		c := string([]rune(m)[:1]) 
		if (c == "." || c == "," || c == "!" || c == "?" || c == ":" || c == ";" || c == "!?"){
			if mots[i-1] == " "{
				result = append(result, "")
			}
		}
		result = append(result, m)
	}
	return result
}