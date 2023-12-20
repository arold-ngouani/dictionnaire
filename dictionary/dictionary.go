package dictionary

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Word struct {
	Mot        string `json:"mot"`
	Definition string `json:"definition"`
}

type Dictionary struct {
	Entries    []Word `json:"entries"`
	addChannel chan Word
	delChannel chan string
}

func New() *Dictionary {
	d := &Dictionary{
		addChannel: make(chan Word),
		delChannel: make(chan string),
	}
	go d.processChannels()
	return d
}

// func (d *Dictionary) Add(mot, definition string) error {
// 	entry := Word{Mot: mot, Definition: definition}
// 	d.Entries = append(d.Entries, entry)
// 	return d.saveToFile()
// }

func (d *Dictionary) Add(mot, definition string) {
	d.addChannel <- Word{Mot: mot, Definition: definition}
}

func (d *Dictionary) Remove(mot string) {
	d.delChannel <- mot
}

// func (d *Dictionary) Get(mot string) (string, error) {
// 	for _, entry := range d.Entries {
// 		if entry.Mot == mot {
// 			return entry.Definition, nil
// 		}
// 	}
// 	return "", fmt.Errorf("Le mot '%s' n'a pas été trouvé dans le dictionnaire", mot)
// }

// func (d *Dictionary) Remove(mot string) error {
// 	for i, entry := range d.Entries {
// 		if entry.Mot == mot {
// 			d.Entries = append(d.Entries[:i], d.Entries[i+1:]...)
// 			return d.saveToFile()
// 		}
// 	}
// 	return fmt.Errorf("Le mot '%s' n'a pas été trouvé dans le dictionnaire", mot)
// }

func (d *Dictionary) List() error {
	for _, entry := range d.Entries {
		fmt.Printf("%s: %s\n", entry.Mot, entry.Definition)
	}
	return nil
}

func (d *Dictionary) saveToFile() error {
	data, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile("dictionary.json", data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func LoadDictionaryFromFile() (*Dictionary, error) {
	fileData, err := ioutil.ReadFile("dictionary.json")
	if err != nil {
		return nil, err
	}

	var dictionary Dictionary
	err = json.Unmarshal(fileData, &dictionary)
	if err != nil {
		return nil, err
	}
	return &dictionary, nil
}

func (d *Dictionary) processChannels() {
	for {
		select {
		case motToAdd := <-d.addChannel:
			d.Entries = append(d.Entries, motToAdd)
			d.saveToFile()
		case motToRemove := <-d.delChannel:
			for i, entry := range d.Entries {
				if entry.Mot == motToRemove {
					d.Entries = append(d.Entries[:i], d.Entries[i+1:]...)
					d.saveToFile()
					break
				}
			}
		}

	}
}
