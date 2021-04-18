package gunmodule

import "github.com/dm1trypon/engine-srv/goapi/models/commondata"

/*
GunModule - data of weapon module object.
	- Data  - contains:
		- kind <string> - kind of module.
		- position <commondata.Position> - data of gun module position.
		- size <commondata.Size> - data of gun module size.
		- chance <int> - probability of gun module.
*/
type GunModule struct {
	kind     string
	chance   int
	position commondata.Position
	size     commondata.Size
}
