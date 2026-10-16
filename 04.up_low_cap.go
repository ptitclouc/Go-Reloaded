package main

import (
	"fmt"
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
	