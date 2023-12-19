package dictionary

import (
	"fmt"
)

type Dictionary map[string]string

func New() Dictionary {
	return Dictionary{}
}

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

	for _, mot := range mots {
		fmt.Printf("%s: %s\n", mot, d[mot])
	}
}
