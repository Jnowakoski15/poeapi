package poeapi

//Socket for POE items
type Socket struct {
	Group   int    `json:"group"`
	Attr    string `json:"attr"`
	SColour string `json:"sColour"`
}

//Category of poe items
type Category struct {
	Accessories []string `json:"accessories"`
	Armour      []string `json:"armour"`
	Jewels      []string `json:"jewels"`
	Weapons     []string `json:"weapons"`
}

// Property values supporting POE Items
type Property struct {
	Name        string      `json:"name"`
	Values      interface{} `json:"values"`
	DisplayMode int         `json:"displayMode"`
	Type        int         `json:"type"`
	Progress    float32     `json:"progress"`
}

//Item has all the poe fields
type Item struct {
	AbyssJewel            bool       `json:"abyssJewel"`
	AdditionalProperties  []Property `json:"additionalProperties"`
	ArtFileName           string     `json:"artFilename"`
	Category              []Category `json:"category"`
	Corrupted             bool       `json:"corrupted"`
	CosmeticMods          []string   `json:"cosmeticMods"`
	CraftedMods           []string   `json:"craftedMods"`
	DescrText             string     `json:"descrText"`
	Duplicated            bool       `json:"duplicated"`
	Elder                 bool       `json:"elder"`
	EnchantMods           []string   `json:"enchantMods"`
	ExplicitMods          []string   `json:"explicitMods"`
	FlavourText           []string   `json:"flavourText"`
	FrameType             int        `json:"frameType"`
	Height                int        `json:"h"`
	Icon                  string     `json:"icon"`
	ID                    string     `json:"id"`
	Identified            bool       `json:"identified"`
	Ilvl                  int        `json:"ilvl"`
	ImplicitMods          []string   `json:"implicitMods"`
	InventoryID           string     `json:"inventoryId"`
	IsRelic               bool       `json:"isRelic"`
	League                string     `json:"league"`
	LockedToCharacter     bool       `json:"lockedToCharacter"`
	MaxStackSize          int        `json:"maxStackSize"`
	Name                  string     `json:"name"`
	NextLevelRequirements []Property `json:"nextLevelRequirements"`
	Note                  string     `json:"note"`
	Properties            []Property `json:"properties"`
	ProphecyDiffText      string     `json:"prophecyDiffText"`
	ProphecyText          string     `json:"prophecyText"`
	Requirements          []Property `json:"requirements"`
	SecDescrText          string     `json:"secDescrText"`
	Shaper                bool       `json:"shaper"`
	SocketedItems         []Property `json:"socketedItems"`
	Sockets               []Socket   `json:"sockets"`
	StackSize             int        `json:"stackSize"`
	Support               bool       `json:"support"`
	TalismanTier          int        `json:"talismanTier"`
	TypeLine              string     `json:"typeLine"`
	UtilityMods           []string   `json:"utilityMods"`
	Verified              bool       `json:"verified"`
	Width                 int        `json:"w"`
	X                     int        `json:"x"`
	Y                     int        `json:"y"`
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
