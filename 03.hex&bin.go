package main

import (

	"strconv"
)


func hexbin(mots []string) []string {
	var result []string
	for _, m := range mots {
		if m == "(hex)" || m == "(bin)" {
			if len(result) == 0 {
				continue
			}
			base := 16
			if m == "(bin)" {
				base = 2
			}
			last := len(result) - 1
			n, err := strconv.ParseInt(result[last], base, 64)
			if err == nil {
				result[last] = strconv.FormatInt(n, 10)
			}
			continue
		}
		result = append(result, m)
	}
	return result
}

//LILIAN func 
/* func hexbin(mots []string) []string {
	var result []string
	for i, m := range mots {
		if m == "(hex)" {
			hexconv, err := strconv.ParseInt(mots[i-1], 16, 64)
			last := len(result) - 1
			if err != nil {
				fmt.Println("Erreur de conversion hexadécimale :", err)
				continue
			}
			result[last] = strconv.FormatInt(hexconv, 10)
			continue
		}
		result = append(result, m)
	}
	return result
}
 */