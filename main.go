package main

import (
	"dictionnaire/dictionary"
	"fmt"
	"time"
)

func main() {

	dict := dictionary.New()
	// if err != nil {
	// 	fmt.Println("Erreur lors du chargement du dictionaire : ", err)
	// 	return
	// }

	// dict.Add("v1", "premiere valeur")
	// if err != nil {
	// 	fmt.Println("Erreur lors de l'ajout au dictionnaire :", err)
	// }
	// dict.Add("v2", "seconde valeur")
	// if err != nil {
	// 	fmt.Println("Erreur lors de l'ajout au dictionnaire :", err)
	// }
	// dict.Add("v3", "troisieme valeur")
	// if err != nil {
	// 	fmt.Println("Erreur lors de l'ajout au dictionnaire :", err)
	// }

	go func() {
		dict.Add("v1", "premiere valeur")
		time.Sleep(1 * time.Second)
		// dict.Remove("v1")
	}()

	go func() {
		dict.Add("v2", "seconde valeur")
		time.Sleep(2 * time.Second)
		// dict.Remove("v2")
	}()

	go func() {
		dict.Add("v3", "troisieme valeur")
		time.Sleep(3 * time.Second)
		dict.Remove("v3")
	}()

	// mot := "v2"
	// definition, err := dict.Get(mot)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Printf("La définition de '%s' est : %s\n", mot, definition)
	// }

	time.Sleep(5 * time.Second)

	// motToRemove := "v3"
	// fmt.Printf("%s a été enlever du dictionaire...\n", motToRemove)
	// err = dict.Remove(motToRemove)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	fmt.Println("\nListe des mots et de leurs définitions: ")
	dict.List()

}
