package cartridge

import (
	"github.com/dm1trypon/engine-srv/goapi/models/commondata"
	"github.com/dm1trypon/engine-srv/goapi/player/weapon/cartridge/effect"
)

/*
Cartridge - contains the characteristics of the weapon.
	- Data - contains:
		- lifeTime <int> - time of life the bullet.
		- damage <int> - damage(health).
		- speed <int> - speed.
		- size <commondata.Size> - data of size.
		- effects <[]BulletEffect> - list of effects.
		- radiusOfDestruction <int> - radius of destruction.
		- triggerRadius <int> - trigger radius.
		- triggerDelay <int> - trigger delay.
		- ricochet <int> - number of ricochet.
*/
type Cartridge struct {
	lifeTime            int
	damage              int
	speed               int
	size                commondata.Size
	effects             []effect.Effect
	radiusOfDestruction int
	triggerRadius       int
	triggerDelay        int
	ricochet            int
}
