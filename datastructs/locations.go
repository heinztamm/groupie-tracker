package datastructs

type LocationsJson struct {
	Indexes []Index `json:"index"`
}

type Index struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
	DatesURL  string   `json:"dates"`
}
