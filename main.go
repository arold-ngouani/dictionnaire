package main

import (
	"fmt"
)

type Dictionary map[string]int

func (d Dictionary) Add(mot string, valeur int) {
	d[mot] = valeur
}

func (d Dictionary) Get(mot string) int {
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

	for _, mot := range mots {
		fmt.Printf("%s: %d\n", mot, d[mot])
	}
}

func (d Dictionary) List2() {

}

func main() {

	dict := make(Dictionary)

	dict.Add("v1", 1)
	dict.Add("v2", 12)
	dict.Add("v3", 123)

	mot := "v2"
	fmt.Printf("Definition of %s: %d\n", mot, dict.Get(mot))

	motToRemove := "v3"
	fmt.Printf("enlever %s du dictionaire...\n", motToRemove)
	dict.Remove(motToRemove)

	fmt.Println("\n Liste du Dictionnaire:")
	dict.List()
}
