package specialrune

import (
	"github.com/dm1trypon/engine-srv/goapi/models/commondata"
	"github.com/dm1trypon/engine-srv/goapi/specialrune/effect"
)

func (s *SpecialRune) Create() {
	s = &SpecialRune{
		effect: effect.Effect{},
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

func (s *SpecialRune) SetEffect(effect effect.Effect) {
	s.effect = effect
}

func (s *SpecialRune) SetPosition(position commondata.Position) {
	s.position = position
}

func (s *SpecialRune) SetSize(size commondata.Size) {
	s.size = size
}

func (s *SpecialRune) GetEffect() effect.Effect {
	return s.effect
}

func (s *SpecialRune) GetPosition() commondata.Position {
	return s.position
}

func (s *SpecialRune) GetSize() commondata.Size {
	return s.size
}
