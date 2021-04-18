package module

import (
	"github.com/dm1trypon/engine-srv/goapi/models/commondata"
	"github.com/dm1trypon/engine-srv/goapi/player/weapon/module/effect"
)

/*
Module - contains the characteristics of the weapon.
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
type Module struct {
	kind              string
	lifeTime          int
	bulletSpeed       int
	reload            int
	damage            int
	bulletSize        commondata.Size
	cartridgeCapacity int
	rateOfFire        int
	isExist           bool
	effects           []effect.Effect
}
