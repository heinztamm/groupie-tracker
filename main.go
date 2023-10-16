package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Response []struct {
	Artist struct {
		Id        int `json:"id"`
		imagelink string
		Name      string `json:"name"`
		Members   []string
	}
}

func main() {
	rootURL := "https://groupietrackers.herokuapp.com/api/artists"
	res, err := http.Get(rootURL)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
	var result Response
	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}
	// fmt.Println(PrettyPrint(result))
	fmt.Println(result[1].Artist)
}

func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
