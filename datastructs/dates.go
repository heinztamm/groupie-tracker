package datastructs

type DatesJson struct {
	Indexes []IndexDates `json:"index"`
}

type IndexDates struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}
