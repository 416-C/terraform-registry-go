package pkg

type Errors struct {
	Errors []string `json:"errors"`
}

type Meta struct {
	Limit         int `json:"limit"`
	CurrentOffset int `json:"current_offset"`
	NextOffset    int `json:"next_offset"`
	PrevOffset    int `json:"prev_offset"`
}
