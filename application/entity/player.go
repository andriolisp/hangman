package entity

//Player struct will hold the user information
type Player struct {
	Num        int  `json:"num"`
	Points     int  `json:"points"`
	Dead       bool `json:"dead"`
	Turn       bool `json:"turn"`
	Tentatives int  `json:"tentatives"`
}
