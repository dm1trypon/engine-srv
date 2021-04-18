package manager

type moveMSG struct {
	Method     string   `json:"method"`
	Nickname   string   `json:"nickname"`
	Directions []string `json:"directions"`
}

type choiceWeapon struct {
	Method   string `json:"method"`
	Nickname string `json:"nickname"`
	Slot     string `json:"slot"`
}

type mouseMSG struct {
	Method    string   `json:"method"`
	Nickname  string   `json:"nickname"`
	IsPressed bool     `json:"is_pressed"`
	Position  position `json:"position"`
}

type respErr struct {
	Method string `json:"method"`
	ErrMsg string `json:"error"`
}

type position struct {
	X int `json:"x"`
	Y int `json:"y"`
}
