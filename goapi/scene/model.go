package scene

import "github.com/dm1trypon/engine-srv/goapi/models/commondata"

/*
Scene - object within which game events take place.
	- Data - contains:
		- Size <commondata.Size> - data of scene size.
*/
type Scene struct {
	size     commondata.Size
	position commondata.Position
}
