package main

import (
	"dictionnaire/dictionary"
	"fmt"
)

func main() {

	dict, err := dictionary.LoadDictionaryFromFile()
	if err != nil {
		fmt.Println("Erreur lors du chargement du dictionaire : ", err)
		return
	}

	dict.Add("v1", "premiere valeur")
	if err != nil {
		fmt.Println("Erreur lors de l'ajout au dictionnaire :", err)
	}
	dict.Add("v2", "seconde valeur")
	if err != nil {
		fmt.Println("Erreur lors de l'ajout au dictionnaire :", err)
	}
	dict.Add("v3", "troisieme valeur")
	if err != nil {
		fmt.Println("Erreur lors de l'ajout au dictionnaire :", err)
	}

	mot := "v2"
	definition, err := dict.Get(mot)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("La définition de '%s' est : %s\n", mot, definition)
	}

	motToRemove := "v3"
	fmt.Printf("%s a été enlever du dictionaire...\n", motToRemove)
	err = dict.Remove(motToRemove)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("\nListe des mots et de leurs définitions: ")
	err = dict.List()
	if err != nil {
		fmt.Println("Erreur lors de l'affichage du dictionnaire :", err)
	}
}
