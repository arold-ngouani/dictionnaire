package dictionary

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Word struct {
	Mot        string `json:"mot"`
	Definition string `json:"definition"`
}

type Dictionary struct {
	Entries    []Word `json:"entries"`
	addChannel chan Word
	delChannel chan string
	done       chan struct{} // Channel pour signaler la fin du traitement
}

func New() *Dictionary {
	d := &Dictionary{
		addChannel: make(chan Word),
		delChannel: make(chan string),
		done:       make(chan struct{}),
	}
	go d.processChannels()
	return d
}

func (d *Dictionary) Add(mot, definition string) {
	d.addChannel <- Word{Mot: mot, Definition: definition}
}

func (d *Dictionary) Remove(mot string) {
	d.delChannel <- mot
}

func (d *Dictionary) Get(mot string) (string, error) {
	for _, entry := range d.Entries {
		if entry.Mot == mot {
			return entry.Definition, nil
		}
	}
	return "", fmt.Errorf("Le mot '%s' n'a pas été trouvé dans le dictionnaire", mot)
}

func (d *Dictionary) List() []Word {
	return d.Entries
}

func (d *Dictionary) SaveToFile() error {
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
		case mot := <-d.addChannel:
			d.Entries = append(d.Entries, mot)
			d.SaveToFile()
		case mot := <-d.delChannel:
			for i, entry := range d.Entries {
				if entry.Mot == mot {
					d.Entries = append(d.Entries[:i], d.Entries[i+1:]...)
					break
				}
			}
			d.SaveToFile()
		case <-d.done:
			return
		}
	}
}

func (d *Dictionary) AddHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var word Word
	err := json.NewDecoder(r.Body).Decode(&word)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	d.Add(word.Mot, word.Definition)
	w.WriteHeader(http.StatusCreated)
}

func (d *Dictionary) GetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	mot := r.URL.Query().Get("mot")
	definition, err := d.Get(mot)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	response := struct {
		Mot        string `json:"mot"`
		Definition string `json:"definition"`
	}{Mot: mot, Definition: definition}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (d *Dictionary) RemoveHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	mot := r.URL.Query().Get("mot")
	d.Remove(mot)
	w.WriteHeader(http.StatusOK)
}

func (d *Dictionary) ListHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	entries := d.List()
	response := struct {
		Entries []Word `json:"entries"`
	}{Entries: entries}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
