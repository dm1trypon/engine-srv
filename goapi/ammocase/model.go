package ammocase

import "github.com/dm1trypon/engine-srv/goapi/models/commondata"

/*
AmmoCase - data of ammunition chest.
	- Data - contains:
		- kind <string> - kind of case.
		- position <commondata.Position> - data of ammo case position.
		- size <commondata.Size> - data of ammo case size.
		- health <int> - object health.
		- lootList <[]string> - list of supports loots.
		- chance <int> - probability of ammo case.
*/
type AmmoCase struct {
	health   int
	kind     string
	position commondata.Position
	size     commondata.Size
	chance   int
	lootList []string
}
