package main


func apostrophes(mots []string) []string {
	var result []string
	count := 0
	for _, m := range mots {
		c := string([]rune(m)[:1])
		d := string([]rune(m)[:len(m)-1])
		if (c == `'`) && (d == `'`) && count%2 != 0 {
			count++
			result[len(result)-1] += m
			continue
		}else if (c == `'`) && (d == `'`) && count%2 == 0{
			count++
			result[len(result)+1] += m
			continue
		}
		result = append(result, m)
	}
	return result
}
