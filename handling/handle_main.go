package handling

import (
	"html/template"
	"net/http"

	"01.kood.tech/git/kartamm/groupie-tracker/JsonParse"
)

func HandleMain(w http.ResponseWriter, r *http.Request) {
	artistsURL := "https://groupietrackers.herokuapp.com/api/artists"
	artists := JsonParse.FetchArtistsStruct(artistsURL)
	// fs := http.FileServer(http.Dir("../templates"))
	// http.Handle("/templates/", http.StripPrefix("/templates", fs))
	tmpl, err := template.ParseFiles("templates/home.html")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(w, artists)
	if err != nil {
		panic(err)
	}
}