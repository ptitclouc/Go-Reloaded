package main

import (
	"fmt"
	"os"
)

func Entree() ([]byte, bool) {
	contenu, err := os.ReadFile(os.Args[1])
	if err != nil {
		return nil, false
	}
	return contenu, true
}

func Sortie(texte string) {
	err := os.WriteFile(os.Args[2], []byte(texte), 0644)
	if err != nil {
		fmt.Println("Erreur d'écriture")
	}
}