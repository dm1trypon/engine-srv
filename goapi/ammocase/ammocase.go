package ammocase

import "github.com/dm1trypon/engine-srv/goapi/models/commondata"

func (a *AmmoCase) Create() *AmmoCase {
	a = &AmmoCase{
		health: 0,
		kind:   "",
		position: commondata.Position{
			X: 0,
			Y: 0,
		},
		size: commondata.Size{
			Width:  0,
			Height: 0,
		},
		chance:   0,
		lootList: []string{},
	}

	return a
}

func (a *AmmoCase) SetHealth(health int) {
	a.health = health
}

func (a *AmmoCase) SetKind(kind string) {
	a.kind = kind
}

func (a *AmmoCase) SetPosition(position commondata.Position) {
	a.position = position
}

func (a *AmmoCase) SetSize(size commondata.Size) {
	a.size = size
}

func (a *AmmoCase) SetChance(chance int) {
	a.chance = chance
}

func (a *AmmoCase) SetLootList(lootList []string) {
	a.lootList = lootList
}

func (a *AmmoCase) GetHealth() int {
	return a.health
}

func (a *AmmoCase) GetKind() string {
	return a.kind
}

func (a *AmmoCase) GetPosition() commondata.Position {
	return a.position
}

func (a *AmmoCase) GetSize() commondata.Size {
	return a.size
}

func (a *AmmoCase) GetChance() int {
	return a.chance
}

func (a *AmmoCase) GetLootList() []string {
	return a.lootList
}
