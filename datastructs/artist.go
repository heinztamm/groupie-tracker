package datastructs

type Artist struct {
	Id              int      `json:"id"`
	ImageURL        string   `json:"image,omitempty"`
	Name            string   `json:"name"`
	Members         []string `json:"members"`
	CreationDate    int      `json:"creationDate"`
	FirstAlbum      string   `json:"firstAlbum"`
	LocationsURL    string   `json:"locations"`
	ConcertDatesURL string   `json:"concertDates"`
	RelationsURL    string   `json:"relations"`
	SearchResult    []Artist
}
