package entity

//Detail will provide the information of each play
type Detail struct {
	Player     int    `json:"player"`
	Points     int    `json:"points"`
	Letter     string `json:"letter"`
	Found      bool   `json:"found"`
	Sequential int    `json:"sequential"`
}

//Details is a collection of Detail struct
type Details []Detail

//HasLetter will check if the array has an struct with the same letter
func (d Details) HasLetter(letter string) bool {
	for _, c := range d {
		if c.Letter == letter {
			return true
		}
	}

	return false
}
