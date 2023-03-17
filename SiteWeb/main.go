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
	Actor    string `json:"actor"`
	Image    string `json:"image"`
}

type Spells []struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CharacterInfo struct {
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
	http.HandleFunc("/gryffindor", gryffindor)
	http.HandleFunc("/slytherin", slytherin)
	http.HandleFunc("/ravenclaw", ravenclaw)
	http.HandleFunc("/hufflepuff", hufflepuff)
	http.HandleFunc("/harry", harryPotter)
	http.HandleFunc("/hermione", hermione)
	http.HandleFunc("/ron", ron)
	http.HandleFunc("/show-character", showCharacter)

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

func gryffindor(w http.ResponseWriter, r *http.Request) {
	tmpl4, err := template.ParseFiles("HTML/gryffindor.html")
	if err != nil {
		panic(err)
	}
	url := "https://hp-api.onrender.com/api/characters/house/Gryffindor"
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var characters []CharacterInfo
	err = json.NewDecoder(resp.Body).Decode(&characters)
	if err != nil {
		panic(err)
	}

	err = tmpl4.Execute(w, characters)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func slytherin(w http.ResponseWriter, r *http.Request) {
	tmpl5, err := template.ParseFiles("HTML/slytherin.html")
	if err != nil {
		panic(err)
	}
	url := "https://hp-api.onrender.com/api/characters/house/slytherin"
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var characters []CharacterInfo
	err = json.NewDecoder(resp.Body).Decode(&characters)
	if err != nil {
		panic(err)
	}

	err = tmpl5.Execute(w, characters)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func ravenclaw(w http.ResponseWriter, r *http.Request) {
	tmpl6, err := template.ParseFiles("HTML/ravenclaw.html")
	if err != nil {
		panic(err)
	}
	url := "https://hp-api.onrender.com/api/characters/house/ravenclaw"
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var characters []CharacterInfo
	err = json.NewDecoder(resp.Body).Decode(&characters)
	if err != nil {
		panic(err)
	}

	err = tmpl6.Execute(w, characters)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func hufflepuff(w http.ResponseWriter, r *http.Request) {
	tmpl7, err := template.ParseFiles("HTML/hufflepuff.html")
	if err != nil {
		panic(err)
	}
	url := "https://hp-api.onrender.com/api/characters/house/hufflepuff"
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var characters []CharacterInfo
	err = json.NewDecoder(resp.Body).Decode(&characters)
	if err != nil {
		panic(err)
	}

	err = tmpl7.Execute(w, characters)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func showCharacter(w http.ResponseWriter, r *http.Request) {
	selectedCharacter := r.FormValue("character")

	url := "https://hp-api.onrender.com/api/characters"
	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var characters []CharacterInfo
	err = json.NewDecoder(resp.Body).Decode(&characters)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var selectedCharacterInfo CharacterInfo
	for _, character := range characters {
		if character.Name == selectedCharacter {
			selectedCharacterInfo = character
			break
		}
	}
	tmpl, err := template.ParseFiles("HTML/show-character.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, selectedCharacterInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func harryPotter(w http.ResponseWriter, r *http.Request) {
	url := "https://hp-api.onrender.com/api/characters"
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var characters []CharacterInfo
	err = json.NewDecoder(resp.Body).Decode(&characters)
	if err != nil {
		panic(err)
	}

	var harry CharacterInfo
	for _, character := range characters {
		if character.Name == "Harry Potter" {
			harry = character
			break
		}
	}

	tmpl, err := template.ParseFiles("HTML/harry.html")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(w, harry)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func hermione(w http.ResponseWriter, r *http.Request) {
	url := "https://hp-api.onrender.com/api/characters"
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var characters []CharacterInfo
	err = json.NewDecoder(resp.Body).Decode(&characters)
	if err != nil {
		panic(err)
	}

	var hermione CharacterInfo
	for _, character := range characters {
		if character.Name == "Hermione Granger" {
			hermione = character
			break
		}
	}

	tmpl, err := template.ParseFiles("HTML/hermione.html")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(w, hermione)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func ron(w http.ResponseWriter, r *http.Request) {
	url := "https://hp-api.onrender.com/api/characters"
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var characters []CharacterInfo
	err = json.NewDecoder(resp.Body).Decode(&characters)
	if err != nil {
		panic(err)
	}

	var ron CharacterInfo
	for _, character := range characters {
		if character.Name == "Ron Weasley" {
			ron = character
			break
		}
	}

	tmpl, err := template.ParseFiles("HTML/ron.html")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(w, ron)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
