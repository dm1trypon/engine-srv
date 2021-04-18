package gunmodule

import "github.com/dm1trypon/engine-srv/goapi/models/commondata"

func (g *GunModule) Create() {
	g = &GunModule{
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

func (g *GunModule) GetKind() string {
	return g.kind
}

func (g *GunModule) GetChance() int {
	return g.chance
}

func (g *GunModule) GetPosition() commondata.Position {
	return g.position
}

func (g *GunModule) GetSize() commondata.Size {
	return g.size
}

func (g *GunModule) SetKind(kind string) {
	g.kind = kind
}

func (g *GunModule) SetChance(chance int) {
	g.chance = chance
}

func (g *GunModule) SetPosition(position commondata.Position) {
	g.position = position
}

func (g *GunModule) SetSize(size commondata.Size) {
	g.size = size
}
