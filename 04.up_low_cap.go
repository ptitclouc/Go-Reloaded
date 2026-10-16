package main

import (
	"strings"
)

func Up(mots []string) []string{
	var result []string
	for _, m := range mots {
		if m == "(up)"{
			if len(result) == 0 {
				continue
			}
			last := len(result) - 1
			result[last] = strings.ToUpper(result[last])
			continue
		}
		result = append(result, m)
	} 
	return result
}

func Low(mots []string) []string{
	var result []string
	for _, m := range mots {
		if m == "(low)"{
			if len(result) == 0 {
				continue
			}
			last := len(result) - 1
			result[last] = strings.ToLower(result[last])
			continue
		}
		result = append(result, m)
	} 
	return result
}

func Cap(mots []string) []string{
	var result []string
	for _, m := range mots {
		if m == "(cap)"{
			if len(result) == 0 {
				continue
			}
			last := len(result) - 1
			result[last] = Capitalize(result[last])
			continue
		}

		result = append(result, m)
	} 
	return result
}

/* func Cap(mots []string) []string{
	var result []string
	n := 0
	for _, m := range mots {
		for i := 0; i < n; i++ {
			if (m == "(cap)" || m == "(cap, "n")"){
				if len(result) == 0 {
					continue
				}
				last := len(result) - 1
				result[last] = Capitalize(result[last])
				continue
			}
		result = append(result, m)
		}
	} 
	return result
} */