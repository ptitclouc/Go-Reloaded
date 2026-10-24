package main

import (
	"strings"
	"unicode"
)

func ponctuation(mots []string) []string {
	var result []string
	for i, m := range mots {
		c := string([]rune(m)[:1])
		if (c == "." || c == "," || c == "!" || c == "?" || c == ":" || c == ";") && i > 0 {
			result[len(result)-1] = strings.TrimRightFunc(result[len(result)-1], unicode.IsSpace) 
			continue
		}
		result = append(result, m)
	}
	return result
}