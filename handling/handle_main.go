package handling

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"01.kood.tech/git/kartamm/groupie-tracker/JsonParse"
	"01.kood.tech/git/kartamm/groupie-tracker/datastructs"
)

type TemplateData struct {
	Artists    []datastructs.Artist
	SearchData Search
}
type Search struct {
	SearchResult []datastructs.Artist
	Input        string
}

func HandleMain(w http.ResponseWriter, r *http.Request) {
	artistsURL := "https://groupietrackers.herokuapp.com/api/artists"
	data := TemplateData{}
	data.Artists = JsonParse.FetchArtistsStruct(artistsURL)
	tmpl, err := template.ParseFiles("templates/home.html")
	if err != nil {
		panic(err)
	}
	switch r.Method {
	case "GET":
		err = tmpl.Execute(w, data)
		if err != nil {
			panic(err)
		}
	case "POST":
		searchInput := r.FormValue("input")
		data.SearchData.Input = searchInput
		data.SearchData.SearchResult = []datastructs.Artist{}
		fmt.Println(searchInput)
		for _, obj := range data.Artists {
			if strings.Contains(obj.Name, data.SearchData.Input) {
				data.SearchData.SearchResult = append(data.SearchData.SearchResult, obj)
				fmt.Println(obj.Name)
			}
		}
		err = tmpl.Execute(w, data)
		if err != nil {
			panic(err)
		}
	}
}
