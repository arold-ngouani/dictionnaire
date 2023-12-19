package main

import (
	"dictionnaire/dictionary"
	"fmt"
)

func main() {

	dict := dictionary.New()

	dict.Add("v1", "premiere valeur")
	dict.Add("v2", "seconde valeur")
	dict.Add("v3", "troisieme valeur")

	mot := "v2"
	definition := dict.Get(mot)
	if definition != "" {
		fmt.Printf("La définition de '%s' est : %s\n", mot, definition)
	} else {
		fmt.Printf("Le mot '%s' n'a pas été trouvé dans le dictionnaire.\n", mot)
	}
	fmt.Printf("Definition de %s: %s\n", mot, dict.Get(mot))

	motToRemove := "v3"
	fmt.Printf("%s a été enlever du dictionaire...\n", motToRemove)
	dict.Remove(motToRemove)

	fmt.Println("\nListe des mots et de leurs définitions: ")
	dict.List()
}
