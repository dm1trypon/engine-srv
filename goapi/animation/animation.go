package animation

import "github.com/dm1trypon/engine-srv/goapi/models/commondata"

func (a *Animation) Create() {
	a = &Animation{
		Kind:     "",
		Duration: 0,
		Position: commondata.Position{
			X: 0,
			Y: 0,
		},
		Size: commondata.Size{
			Width:  0,
			Height: 0,
		},
	}
}
