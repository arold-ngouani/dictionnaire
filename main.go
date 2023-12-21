package main

import (
	"dictionnaire/dictionary"
	"fmt"
	"net/http"
	"os"
)

func main() {
	// dict, err := dictionary.LoadDictionaryFromFile()
	// if err != nil {
	// 	dict = dictionary.New()
	// 	fmt.Println("Le fichier dictionary.json n'existe pas ou est invalide. Création d'un nouveau dictionnaire.")
	// }
	_, err := os.Stat("dictionary.json")
    var dict *dictionary.Dictionary

    if os.IsNotExist(err) {
        // Le fichier dictionary.json n'existe pas, créer un nouveau dictionnaire
        dict = dictionary.New()
        fmt.Println("Le fichier dictionary.json n'existe pas. Création d'un nouveau dictionnaire.")

        // Enregistrez le nouveau dictionnaire dans dictionary.json
        if err := dict.SaveToFile(); err != nil {
            fmt.Println("Erreur lors de la création du fichier dictionary.json:", err)
            return
        }
    } else {
        // Le fichier dictionary.json existe, chargez les données
        dict, err = dictionary.LoadDictionaryFromFile()
        if err != nil {
            // Gérer l'erreur de chargement du fichier
            dict = dictionary.New()
            fmt.Println("Le fichier dictionary.json est invalide. Création d'un nouveau dictionnaire.")
        }
    }

    http.HandleFunc("/add", dict.AddHandler)
    http.HandleFunc("/get", dict.GetHandler)
    http.HandleFunc("/remove", dict.RemoveHandler)
    http.HandleFunc("/list", dict.ListHandler)

    fmt.Println("Serveur démarré sur le port 8080")
    http.ListenAndServe(":8080", nil)

}
