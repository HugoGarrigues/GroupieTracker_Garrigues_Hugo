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

type Spells []struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Wand struct {
	Wood   string `json:"wood"`
	Core   string `json:"core"`
	Length int    `json:"length"`
}

type CharacterInfo []struct {
	Name            string   `json:"name"`
	AlternateNames  []string `json:"alternate_names"`
	Species         string   `json:"species"`
	Gender          string   `json:"gender"`
	House           string   `json:"house"`
	DateOfBirth     string   `json:"dateOfBirth"`
	YearOfBirth     int      `json:"yearOfBirth"`
	Wizard          bool     `json:"wizard"`
	Ancestry        string   `json:"ancestry"`
	EyeColour       string   `json:"eyeColour"`
	HairColour      string   `json:"hairColour"`
	Wand            Wand     `json:"wand"`
	Patronus        string   `json:"patronus"`
	HogwartsStudent bool     `json:"hogwartsStudent"`
	HogwartsStaff   bool     `json:"hogwartsStaff"`
	Actor           string   `json:"actor"`
	AlternateActors []any    `json:"alternate_actors"`
	Alive           bool     `json:"alive"`
	Image           string   `json:"image"`
}

func main() {
	static := http.FileServer(http.Dir("CSS"))
	http.Handle("/CSS/", http.StripPrefix("/CSS/", static))

	http.HandleFunc("/", index)
	http.HandleFunc("/characters", characters)
	http.HandleFunc("/spells", spells)

	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	tmpl1, err := template.ParseFiles("HTML/index.html")
	if err != nil {
		panic(err)
	}
	tmpl1.Execute(w, nil)
}

func characters(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("HTML/characters.html")
	if err != nil {
		panic(err)
	}

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

func spells(w http.ResponseWriter, r *http.Request) {
	tmpl3, err := template.ParseFiles("HTML/spells.html")
	if err != nil {
		panic(err)
	}
	url := "https://hp-api.onrender.com/api/spells"
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var spells Spells
	err = json.NewDecoder(resp.Body).Decode(&spells)
	if err != nil {
		panic(err)
	}

	err = tmpl3.Execute(w, spells)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
