package handling

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"

	"01.kood.tech/git/kartamm/groupie-tracker/JsonParse"
)

func HandleDetails(w http.ResponseWriter, r *http.Request) {
	prefix := "/details/"
	IDstr := strings.TrimPrefix(r.URL.Path, prefix)
	ID, _ := strconv.Atoi(IDstr)

	data := TemplateData{}

	artistsURL := "https://groupietrackers.herokuapp.com/api/artists"
	data.Artists = JsonParse.FetchArtistsStruct(artistsURL)
	selectedArtist := data.Artists[ID-1]
	fmt.Println(selectedArtist.Name)
	tmpl, err := template.ParseFiles("templates/details.html")
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		panic(err)
	}
}
