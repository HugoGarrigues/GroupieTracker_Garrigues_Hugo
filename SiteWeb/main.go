package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"text/template"
)

type Characters struct {
	Name     string `json:"name"`
	House    string `json:"house"`
	Patronus string `json:"patronus"`
	Alive    bool   `json:"alive"`
	Image    string `json:"image"`
}

func main() {
	// Lier le fichier css qui est dans ../CSS/style.css //
	static := http.FileServer(http.Dir("CSS"))
	http.Handle("/CSS/", http.StripPrefix("/CSS/", static))

	tmpl, err := template.ParseFiles("HTML/index.html")
	if err != nil {
		panic(err)
	}
	// Création du serveur //
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			err := r.ParseForm()
			if err != nil {
				http.Error(w, "Bad Request", http.StatusBadRequest)
				return
			}

			class := r.FormValue("class")
			if err != nil {
				http.Error(w, "NUL", http.StatusBadRequest)
				return
			} else if class == "Gryffindor" {
				class = "Gryffindor"
			} else if class == "Slytherin" {
				class = "Slytherin"
			} else if class == "Ravenclaw" {
				class = "Ravenclaw"
			} else if class == "Hufflepuff" {
				class = "Hufflepuff"
			}

			characters := characters_class(class)

			err = tmpl.Execute(w, characters)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		} else {
			tmpl.Execute(w, nil)
		}
	})

	http.ListenAndServe(":8080", nil)
}

func characters_class(class string) []Characters {
	url := fmt.Sprintf("https://hp-api.onrender.com/api/characters/house/%s", class)
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var characters []Characters
	err = json.NewDecoder(resp.Body).Decode(&characters)
	if err != nil {
		panic(err)
	}

	return characters
}
