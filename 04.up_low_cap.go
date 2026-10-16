package main

import (

	"strings"
)

func up(mots []string) []string{
	var result []string
	for _, m := range mots {
		if m == "(up)"{
			if len(result) == 0 {
				continue
			}
			last := len(result) - 1
			strings.ToUpper(result[last])
			continue
		}
		result = append(result, m)
	} 
	return result
}
	