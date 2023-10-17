package JsonParse

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"01.kood.tech/git/kartamm/groupie-tracker/datastructs"
)

func FetchRelations(URL string) datastructs.RelationsJson {
	res, err := http.Get(URL)
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
	var relations datastructs.RelationsJson
	if err := json.Unmarshal(Data, &relations); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}
	return relations
}
