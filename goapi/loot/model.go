package loot

import "github.com/dm1trypon/engine-srv/goapi/models/commondata"

/*
Loot - data for object that contains ammo, first aid kits and armor.
	- Data - contains:
		- kind <string> - kind of loot.
		- position <commondata.Position> - data of loot position.
		- size <commondata.Size> - data of loot size.
		- chance <int> - probability of loot.
*/
type Loot struct {
	kind     string
	value    int
	chance   int
	position commondata.Position
	size     commondata.Size
}
