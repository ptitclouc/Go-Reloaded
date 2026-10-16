package main

import "strings"

func main() {
	contenu, ok := Entree()
	if !ok {
		return
	}
	mots := decoupage(contenu)
	mots = hexbin(mots)
	mots = Up(mots)
	mots = Low(mots)
	mots = Cap(mots)
	texte := strings.Join(mots, " ")
	Sortie(texte)
}