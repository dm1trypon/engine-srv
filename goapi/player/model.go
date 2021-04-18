package player

import (
	"github.com/dm1trypon/engine-srv/goapi/models/commondata"
	"github.com/dm1trypon/engine-srv/goapi/player/cursor"
	"github.com/dm1trypon/engine-srv/goapi/player/effect"
	"github.com/dm1trypon/engine-srv/goapi/player/weapon"
)

/*
Player - data of characteristics objects of the Players.
	- Data - contains:
		- nickname <string> - nickname.
		- curHealth <int> - current health.
		- maxHealth <int> - limit of health.
		- curArmor <int> - current armor.
		- maxArmor <int> - limit of armor.
		- directions <[]string> - list of current directions.
		- cursor <PlayerCursor> - data of player cursor on screen.
		- position <commondata.Position> - data of player position.
		- speed <commondata.Speed> - data of player speed.
		- size <commondata.Size> - data of player size.
		- weapons <[]PlayerWeapon> - list of weapons data(weapon slots).
		- effects <[]PlayerEffect> - list of current effects data.
		- isRemoved <bool> - event on remove object.
*/
type Player struct {
	nickname   string
	curHealth  int
	maxHealth  int
	curArmor   int
	maxArmor   int
	directions []string
	rotation   int
	cursor     cursor.Cursor
	position   commondata.Position
	speed      commondata.Speed
	size       commondata.Size
	weapons    []weapon.Weapon
	effects    []effect.Effect
	isRemoved  bool
}
