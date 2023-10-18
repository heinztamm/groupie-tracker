package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"01.kood.tech/git/kartamm/groupie-tracker/handling"
)

func main() {

	fs := http.FileServer(http.Dir("./styling"))
	http.Handle("/styling/", http.StripPrefix("/styling", fs))
	http.HandleFunc("/", handling.HandleMain)
	// http.HandleFunc("/details", handling.HandleDetails)
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
