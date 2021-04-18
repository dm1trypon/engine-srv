package bullet

import (
	"github.com/dm1trypon/engine-srv/goapi/bullet/effect"
	"github.com/dm1trypon/engine-srv/goapi/models/commondata"
)

/*
Bullet - data of characteristics objects of the Players.
	- Data - contains:
		- nickname <string> - player's nickname.
		- health <int> - current health.
		- kind <int> - weapon's kind.
		- targetPosition <commondata.Position> - position of player's clicked cursor.
		- position <commondata.Position> - bullet's position.
		- speed <commondata.Speed> - bullet's speed.
		- size <commondata.Size> - bullet's size.
		- lifeTime <int> - bullet's life time.
		- effects <[]effect> - list of current effects data.
		- isRemoved <bool> - event on remove object.
*/
type Bullet struct {
	nickname            string
	health              int
	kind                string
	position            commondata.Position
	targetPosition      commondata.Position
	speed               commondata.Speed
	size                commondata.Size
	radiusOfDestruction int
	rotation            int
	ricochet            int
	triggerRadius       int
	triggerDelay        int
	lifeTime            int
	isRemoved           bool
	effects             []effect.Effect
}
