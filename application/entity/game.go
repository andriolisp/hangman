package entity

//Game will return the information about the current game
type Game struct {
	ID         string             `json:"id"`
	Word       string             `json:"word,omitempty"`
	WordSize   int                `json:"size"`
	Remaining  int                `json:"remaining"`
	Turn       int                `json:"turn"`
	PlayersNum int                `json:"num_players"`
	Winner     int                `json:"winner"`
	Message    string             `json:"message,omitempty"`
	Details    Details            `json:"details,omitempty"`
	Replacers  []string           `json:"replacers"`
	Players    map[string]*Player `json:"players"`
}
