package block

import "github.com/dm1trypon/engine-srv/goapi/models/commondata"

func (b *Block) Create(nickname string) {
	b = &Block{
		kind:     "",
		health:   0,
		position: commondata.Position{},
		size:     commondata.Size{},
	}
}

func (b *Block) SetSize(width, height int) {
	b.size.Width = width
	b.size.Height = height
}

func (b *Block) GetPosition() commondata.Position {
	return b.position
}

func (b *Block) GetSize() commondata.Size {
	return b.size
}

func (b *Block) SetPosition(x, y int) {
	b.position.X = x
	b.position.Y = y
}

func (b *Block) GetHealth() int {
	return b.health
}

func (b *Block) SetHealth(health int) {
	b.health = health
}
