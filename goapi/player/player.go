package player

import (
	"github.com/dm1trypon/engine-srv/goapi/models/commondata"
	"github.com/dm1trypon/engine-srv/goapi/player/cursor"
	"github.com/dm1trypon/engine-srv/goapi/player/effect"
	"github.com/dm1trypon/engine-srv/goapi/player/weapon"
)

func (p *Player) Create() *Player {
	p = &Player{
		nickname:   "",
		curHealth:  0,
		maxHealth:  0,
		curArmor:   0,
		maxArmor:   0,
		directions: []string{},
		cursor:     cursor.Cursor{},
		rotation:   0,
		position: commondata.Position{
			X: 0,
			Y: 0,
		},
		speed: commondata.Speed{
			X:   0,
			Y:   0,
			Max: 0,
		},
		size: commondata.Size{
			Width:  0,
			Height: 0,
		},
		weapons:   []weapon.Weapon{},
		effects:   []effect.Effect{},
		isRemoved: false,
	}

	return p
}

func (p *Player) GetNickname() string {
	return p.nickname
}

func (p *Player) GetCursor() cursor.Cursor {
	return p.cursor
}

func (p *Player) GetPosition() commondata.Position {
	return p.position
}

func (p *Player) GetSize() commondata.Size {
	return p.size
}

func (p *Player) GetSpeed() commondata.Speed {
	return p.speed
}

func (p *Player) GetRotation() int {
	return p.rotation
}

func (p *Player) GetHealth() int {
	return p.curHealth
}

func (p *Player) GetArmor() int {
	return p.curArmor
}

func (p *Player) GetWeapons() []weapon.Weapon {
	return p.weapons
}

func (p *Player) GetEffects() []effect.Effect {
	return p.effects
}

func (p *Player) SetNickname(nickname string) {
	p.nickname = nickname
}

func (p *Player) SetPosition(position commondata.Position) {
	p.position = position
}

func (p *Player) SetSpeed(speed commondata.Speed) {
	p.speed = speed
}

func (p *Player) SetRotation(rotation int) {
	p.rotation = rotation
}

func (p *Player) SetHealth(health int) {
	p.curHealth = health
}

func (p *Player) SetArmor(armor int) {
	p.curArmor = armor
}

func (p *Player) SetCursor(cursor cursor.Cursor) {
	p.cursor = cursor
}

func (p *Player) SetWeapons(weapons []weapon.Weapon) {
	p.weapons = weapons
}

func (p *Player) SetSize(size commondata.Size) {
	p.size = size
}

func (p *Player) SetEffects(effects []effect.Effect) {
	p.effects = effects
}
