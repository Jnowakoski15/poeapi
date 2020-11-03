package poeapi

//Item has all the poe fields
type Item struct {
	AbyssJewel bool `json:"abyssJewel"`
}

//Stash contains all the gear in poe
type Stash struct {
	AccountName       string `json:"accountName"`
	LastCharacterName string `json:"lastCharacterName"`
	ID                string `json:"id"`
	Stash             string `json:"stash"`
	StashType         string `json:"stashType"`
	Items             []Item `json:"items"`
	Public            bool   `json:"public"`
	League            string `json:"league"`
}

//Poe top level response obj
type Poe struct {
	NextChangeID string  `json:"next_change_id"`
	Stashes      []Stash `json:"stashes"`
}
