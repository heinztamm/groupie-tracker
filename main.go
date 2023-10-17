package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"01.kood.tech/git/kartamm/groupie-tracker/handling"
)

func main() {

	// artistsURL := "https://groupietrackers.herokuapp.com/api/artists"
	// artists := JsonParse.FetchArtistsStruct(artistsURL)
	// fmt.Println(artists)
	// locationsURL := "https://groupietrackers.herokuapp.com/api/locations"
	// locations := JsonParse.FetchLocationsStruct(locationsURL)
	// datesURL := "https://groupietrackers.herokuapp.com/api/dates"
	// dates := JsonParse.FetchDates(datesURL)
	// relationsURL := "https://groupietrackers.herokuapp.com/api/relation"
	// relationsJson := JsonParse.FetchRelations(relationsURL)
	// fmt.Println(relationsJson.SliceOfRel[2].DatesLocations["london-uk"][0])

	// Next up, display all the artists on a homepage
	// so, first, set up execution of a homepage
	// fs := http.FileServer(http.Dir("./styling"))
	// http.Handle("/styling/", http.StripPrefix("/styling", fs))
	http.HandleFunc("/", handling.HandleMain)
	PORT := ":8080"
	linkText := "Click here to head to the website..."
	linkURL := "http://localhost" + PORT

	// ANSI escape code for hyperlink
	hyperlink := "\033]8;;" + linkURL + "\033\\" + linkText + "\033]8;;\033\\"
	fmt.Println()
	fmt.Println(hyperlink)
	fmt.Println()

	err := http.ListenAndServe(PORT, nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}

}
