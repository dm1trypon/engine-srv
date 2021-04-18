package cursor

import "github.com/dm1trypon/engine-srv/goapi/models/commondata"

/*
Cursor - contains position data and click status of the cursor on screen.
	- Data - contains:
		- position <commondata.Position> - data that contains the coordinates of the player.
		- isPressed <bool> - click status.
*/
type Cursor struct {
	position  commondata.Position
	isPressed bool
}
