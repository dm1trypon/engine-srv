package gunpart

import "github.com/dm1trypon/engine-srv/goapi/models/commondata"

func (g *GunPart) Create() {
	g = &GunPart{
		kind:   "",
		chance: 0,
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

func (g *GunPart) GetKind() string {
	return g.kind
}

func (g *GunPart) GetChance() int {
	return g.chance
}

func (g *GunPart) GetPosition() commondata.Position {
	return g.position
}

func (g *GunPart) GetSize() commondata.Size {
	return g.size
}

func (g *GunPart) SetKind(kind string) {
	g.kind = kind
}

func (g *GunPart) SetChance(chance int) {
	g.chance = chance
}

func (g *GunPart) SetPosition(position commondata.Position) {
	g.position = position
}

func (g *GunPart) SetSize(size commondata.Size) {
	g.size = size
}
