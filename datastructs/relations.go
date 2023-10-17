package datastructs

type RelationsJson struct {
	SliceOfRel []Relations `json:"index"`
}

type Relations struct {
	ID             int
	DatesLocations map[string][]string
}
