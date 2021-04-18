package cursor

import "github.com/dm1trypon/engine-srv/goapi/models/commondata"

func (c *Cursor) Create() *Cursor {
	c = &Cursor{
		position: commondata.Position{
			X: 0,
			Y: 0,
		},
		isPressed: false,
	}

	return c
}

func (c *Cursor) GetPosition() commondata.Position {
	return c.position
}

func (c *Cursor) GetIsPressed() bool {
	return c.isPressed
}

func (c *Cursor) SetPosition(position commondata.Position) {
	c.position = position
}

func (c *Cursor) SetIsPressed(isPressed bool) {
	c.isPressed = isPressed
}
