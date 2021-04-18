package scene

import "github.com/dm1trypon/engine-srv/goapi/models/commondata"

func (s *Scene) Create() {
	s = &Scene{
		size: commondata.Size{
			Width:  0,
			Height: 0,
		},
		position: commondata.Position{
			X: 0,
			Y: 0,
		},
	}
}

func (s *Scene) GetPosition() commondata.Position {
	return s.position
}

func (s *Scene) GetSize() commondata.Size {
	return s.size
}

func (s *Scene) SetPosition(x, y int) {
	s.position.X = x
	s.position.Y = y
}
