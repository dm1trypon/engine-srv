package guncase

import "github.com/dm1trypon/engine-srv/goapi/models/commondata"

/*
GunCase - data of weapon part chests.
	- Data  - contains:
		- kind <string> - kind of effect.
		- position <commondata.Position> - data of gun case position.
		- size <commondata.Size> - data of gun case size.
		- health <int> - object health.
		- moduleKinds <[]string> - list of supports gun modules.
		- chance <int> - probability of gun case.
*/
type GunCase struct {
	health      int
	kind        string
	chance      int
	moduleKinds []string
	position    commondata.Position
	size        commondata.Size
}
