package loot

import (
	"github.com/dm1trypon/engine-srv/goapi/models/commondata"
)

func (l *Loot) Create() {
	l = &Loot{
		kind:   "",
		value:  0,
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

func (l *Loot) GetValue() int {
	return l.value
}

func (l *Loot) GetKind() string {
	return l.kind
}

func (l *Loot) GetPosition() commondata.Position {
	return l.position
}

func (l *Loot) GetSize() commondata.Size {
	return l.size
}

func (l *Loot) SetPosition(x, y int) {
	l.position.X = x
	l.position.Y = y
}
