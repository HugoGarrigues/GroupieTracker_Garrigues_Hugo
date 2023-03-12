package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"text/template"
)

type personnages struct {
	Name  string `json:"name"`
	House string `json:"house"`
}

func main() {
	// Lien vers le fichier CSS //
	static := http.FileServer(http.Dir("CSS"))
	http.Handle("../CSS/style.css", http.StripPrefix("../CSS/style.css", static))
	// Lien vers le fichier html //
	tmpl, err := template.ParseFiles("../HTML/index.html")
	if err != nil {
		panic(err)
	}
	// Création du serveur //
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
	http.ListenAndServe(":8080", nil)
	// Récupération des données //
	resp, err := http.Get("https://hp-api.onrender.com/api/characters")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	var p personnages
	err = json.NewDecoder(resp.Body).Decode(&p)
	if err != nil {
		panic(err)
	}
	fmt.Println(p.Name)
}
