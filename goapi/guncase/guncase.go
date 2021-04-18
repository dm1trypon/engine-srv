package guncase

import "github.com/dm1trypon/engine-srv/goapi/models/commondata"

func (g *GunCase) Create() {
	g = &GunCase{
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

func (g *GunCase) GetHealth() int {
	return g.health
}

func (g *GunCase) GetKind() string {
	return g.kind
}

func (g *GunCase) GetChance() int {
	return g.chance
}

func (g *GunCase) GetModuleKinds() []string {
	return g.moduleKinds
}

func (g *GunCase) GetPosition() commondata.Position {
	return g.position
}

func (g *GunCase) GetSize() commondata.Size {
	return g.size
}

func (g *GunCase) SetHealth(health int) {
	g.health = health
}

func (g *GunCase) SetKind(kind string) {
	g.kind = kind
}

func (g *GunCase) SetChance(chance int) {
	g.chance = chance
}

func (g *GunCase) SetModuleKinds(moduleKinds []string) {
	g.moduleKinds = moduleKinds
}

func (g *GunCase) SetPosition(position commondata.Position) {
	g.position = position
}

func (g *GunCase) SetSize(size commondata.Size) {
	g.size = size
}
