package weapon

import (
	"github.com/dm1trypon/engine-srv/goapi/player/weapon/cartridge"
	"github.com/dm1trypon/engine-srv/goapi/player/weapon/module"
)

/*
Weapon - contains the characteristics of the weapon.
	- Data - contains:
		- Kind <string> - kind of weapon.
		- CurRateOfFire <int> - current rate of fire(timer).
		- MaxRateOfFire <int> - limit rate of fire.
		- CurParts <int> - current number of weapon parts to assemble.
		- MaxParts <int> - limit of weapon parts.
		- CurReloadTime <int> - current weapon reload time(timer).
		- MaxReloadTime <int> - limit of weapon reload time.
		- CurCountCartridges <int> - current number of cartridges in the store.
		- CartridgeCapacity <int> - maximum number of cartridges in the store.
		- CurCartridgeStores <int> - current number of weapon stores.
		- MaxCartridgeStores <int> - maximum number of weapon stores.
		- Spread <int> - spread of bullets.
		- IsActive <bool> - is active weapon.
		- IsDefault <bool> - is default weapon.
		- Modules <[]WeaponModule> - slots of weapon modules.
		- Bullet <WeaponBullet> - data of weapon bullet.
*/
type Weapon struct {
	kind               string
	curRateOfFire      int
	maxRateOfFire      int
	curParts           int
	maxParts           int
	curReloadTime      int
	maxReloadTime      int
	curCountCartridges int
	cartridgeCapacity  int
	curCartridgeStores int
	maxCartridgeStores int
	spread             int
	isActive           bool
	isDefault          bool
	cartridge          cartridge.Cartridge
	modules            []module.Module
}
