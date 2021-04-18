package block

import "github.com/dm1trypon/engine-srv/goapi/models/commondata"

type Block struct {
	kind     string
	health   int
	position commondata.Position
	size     commondata.Size
}
