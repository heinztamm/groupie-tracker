package JsonParse

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"01.kood.tech/git/kartamm/groupie-tracker/datastructs"
)

func FetchArtistsStruct(artistsURL string) []datastructs.Artist {
	res, err := http.Get(artistsURL)
	if err != nil {
		log.Fatal(err)
	}
	Data, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, Data)
	}
	if err != nil {
		log.Fatal(err)
	}
	var artists []datastructs.Artist
	if err := json.Unmarshal(Data, &artists); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}
	return artists
}
