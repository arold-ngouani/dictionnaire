package main

import (
	"fmt"
	"sort"
)

type Dictionary map[string]string

func (d Dictionary) Add(mot, definition string) {
	d[mot] = definition
}

func (d Dictionary) Get(mot string) string {
	return d[mot]
}

func (d Dictionary) Remove(mot string) {
	delete(d, mot)
}

func (d Dictionary) List() {
	var mots []string
	for mot := range d {
		mots = append(mots, mot)
	}

	sort.Strings(mots)

	for _, mot := range mots {
		fmt.Printf("%s: %s\n", mot, d[mot])
	}
}

func main() {

	dict := make(Dictionary)

	dict.Add("valeur1", "essai")
	dict.Add("valeur2", "essai2")
	dict.Add("valeur3", "essai3")

	mot := "valeur2"
	fmt.Printf("Definition of %s: %s\n", mot, dict.Get(mot))

	motToRemove := "valeur3"
	fmt.Printf("enlever %s du dictionnaire...\n", motToRemove)
	dict.Remove(motToRemove)

	fmt.Println("\n Liste du Dictionnaire:")
	dict.List()
}
