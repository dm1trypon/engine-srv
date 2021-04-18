package cartridge

import (
	"github.com/dm1trypon/engine-srv/goapi/models/commondata"
	"github.com/dm1trypon/engine-srv/goapi/player/weapon/cartridge/effect"
)

func (c *Cartridge) Create() *Cartridge {
	c = &Cartridge{
		lifeTime: 0,
		damage:   0,
		speed:    0,
		size: commondata.Size{
			Width:  0,
			Height: 0,
		},
		effects:             []effect.Effect{},
		radiusOfDestruction: 0,
		triggerRadius:       0,
		triggerDelay:        0,
		ricochet:            0,
	}

	return c
}

func (c *Cartridge) GetLifeTime() int {
	return c.lifeTime
}

func (c *Cartridge) GetDamage() int {
	return c.damage
}

func (c *Cartridge) GetSpeed() int {
	return c.speed
}

func (c *Cartridge) GetSize() commondata.Size {
	return c.size
}

func (c *Cartridge) GetEffects() []effect.Effect {
	return c.effects
}

func (c *Cartridge) GetRadiusOfDestruction() int {
	return c.radiusOfDestruction
}

func (c *Cartridge) GetTriggerRadius() int {
	return c.triggerRadius
}

func (c *Cartridge) GetTriggerDelay() int {
	return c.triggerDelay
}

func (c *Cartridge) GetRicochet() int {
	return c.ricochet
}

func (c *Cartridge) SetLifeTime(lifeTime int) {
	c.lifeTime = lifeTime
}

func (c *Cartridge) SetDamage(damage int) {
	c.damage = damage
}

func (c *Cartridge) SetSpeed(speed int) {
	c.speed = speed
}

func (c *Cartridge) SetSize(size commondata.Size) {
	c.size = size
}

func (c *Cartridge) SetEffects(effects []effect.Effect) {
	c.effects = effects
}

func (c *Cartridge) SetRadiusOfDestruction(radiusOfDestruction int) {
	c.radiusOfDestruction = radiusOfDestruction
}

func (c *Cartridge) SetTriggerRadius(triggerRadius int) {
	c.triggerRadius = triggerRadius
}

func (c *Cartridge) SetTriggerDelay(triggerDelay int) {
	c.triggerDelay = triggerDelay
}

func (c *Cartridge) SetRicochet(ricochet int) {
	c.ricochet = ricochet
}
