package dictionary

import (
	"fmt"
)

type Dict struct {
	Mot        string
	Definition string
}

type Dictionary map[string]Dict

func (d Dictionary) Add(mot, definition string) {
	d[mot] = Dict{Mot: mot, Definition: definition}
}

func (d Dictionary) Get(mot string) string {
	return d[mot].Definition
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
		fmt.Printf("%s: %s\n", mot, d[mot].Definition)
	}
}

func (d Dictionary) List2() {

}
