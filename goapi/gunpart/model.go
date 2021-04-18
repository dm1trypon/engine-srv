package gunpart

import "github.com/dm1trypon/engine-srv/goapi/models/commondata"

/*
GunPart - data of weapon part object.
	- Data - contains:
		- kind <string> - kind of weapon.
		- position <commondata.Position> - data of gun part position.
		- size <commondata.Size> - data of gun part size.
		- chance <int> - probability of gun part.
*/
type GunPart struct {
	kind     string
	chance   int
	position commondata.Position
	size     commondata.Size
}
