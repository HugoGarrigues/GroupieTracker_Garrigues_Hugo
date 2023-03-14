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
	// Lien vers le dossier CSS //
	static := http.FileServer(http.Dir("CSS"))
	http.Handle("/CSS/", http.StripPrefix("/CSS/", static))

	// Lien vers le dossier HTML(index.html) soit tmpl1 //
	tmpl1, err := template.ParseFiles("HTML/index.html")
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl1.Execute(w, nil)
			return
		}
		tmpl1.Execute(w, nil)
	})

	// Lien vers le dossier HTML(characters.html) soit tmpl2 //
	tmpl2, err := template.ParseFiles("HTML/characters.html")
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/characters", func(w http.ResponseWriter, r *http.Request) {
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

			err = tmpl2.Execute(w, characters)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		} else {
			tmpl2.Execute(w, nil)
		}
	})

	tmpl3, err := template.ParseFiles("HTML/spells.html")
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/spells", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl3.Execute(w, nil)
			return
		}
		tmpl3.Execute(w, nil)
	})

	tmpl4, err := template.ParseFiles("HTML/protagonists.html")
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/protagonists", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl4.Execute(w, nil)
			return
		}
		tmpl4.Execute(w, nil)
	})

	tmpl5, err := template.ParseFiles("HTML/infos.html")
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/infos", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl5.Execute(w, nil)
			return
		}
		tmpl5.Execute(w, nil)
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
