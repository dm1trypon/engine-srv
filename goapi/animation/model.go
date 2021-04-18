package animation

import "github.com/dm1trypon/engine-srv/goapi/models/commondata"

/*
Animation - data of animation effect.
	- Data - contains:
		- Kind <string> - kind of animation effect.
		- Position <commondata.Position> - data of animation effect position.
		- Size <commondata.Size> - data of animation effect size.
		- Duration <int> - duration of animation effect.
*/
type Animation struct {
	Kind     string
	Duration int
	Position commondata.Position
	Size     commondata.Size
}
