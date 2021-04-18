package gunchest

import "github.com/dm1trypon/engine-srv/goapi/models/commondata"

func (g *GunChest) Create() {
	g = &GunChest{
		health:      0,
		kind:        "",
		chance:      0,
		moduleKinds: []string{},
		position: commondata.Position{
			X: 0,
			Y: 0,
		},
		size: commondata.Size{
			Width:  0,
			Height: 0,
		},
	}
}

func (g *GunChest) GetHealth() int {
	return g.health
}

func (g *GunChest) GetKind() string {
	return g.kind
}

func (g *GunChest) GetChance() int {
	return g.chance
}

func (g *GunChest) GetModuleKinds() []string {
	return g.moduleKinds
}

func (g *GunChest) GetPosition() commondata.Position {
	return g.position
}

func (g *GunChest) GetSize() commondata.Size {
	return g.size
}

func (g *GunChest) SetHealth(health int) {
	g.health = health
}

func (g *GunChest) SetKind(kind string) {
	g.kind = kind
}

func (g *GunChest) SetChance(chance int) {
	g.chance = chance
}

func (g *GunChest) SetModuleKinds(moduleKinds []string) {
	g.moduleKinds = moduleKinds
}

func (g *GunChest) SetPosition(position commondata.Position) {
	g.position = position
}

func (g *GunChest) SetSize(size commondata.Size) {
	g.size = size
}
