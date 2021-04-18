package specialrune

import (
	"github.com/dm1trypon/engine-srv/goapi/models/commondata"
	"github.com/dm1trypon/engine-srv/goapi/specialrune/effect"
)

/*
SpecialRune - object with useful effects.
	- Data - contains:
		- Width <int> - scene width.
		- Height <int> - scene height.
		- Effect <Effect> - effect data for Players and Bullets.
*/
type SpecialRune struct {
	effect   effect.Effect
	position commondata.Position
	size     commondata.Size
}
