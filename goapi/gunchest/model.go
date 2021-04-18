package gunchest

import "github.com/dm1trypon/engine-srv/goapi/models/commondata"

/*
GunChest - data of weapon module chests.
	- Data - contains:
		- kind <string> - kind of effect.
		- position <commondata.Position> - data of gun chest position.
		- size <commondata.Size> - data of gun chest size.
		- health <int> - object health.
		- moduleKinds <[]string> - list of supports gun parts.
		- chance <int> - probability of gun chest.
*/
type GunChest struct {
	health      int
	kind        string
	chance      int
	moduleKinds []string
	position    commondata.Position
	size        commondata.Size
}
