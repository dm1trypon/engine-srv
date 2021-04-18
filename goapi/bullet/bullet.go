package bullet

import (
	"github.com/dm1trypon/engine-srv/goapi/bullet/effect"
	"github.com/dm1trypon/engine-srv/goapi/models/commondata"
)

func (b *Bullet) Create() *Bullet {
	b = &Bullet{
		nickname:            "",
		health:              0,
		kind:                "",
		position:            commondata.Position{X: 0, Y: 0},
		targetPosition:      commondata.Position{X: 0, Y: 0},
		speed:               commondata.Speed{X: 0, Y: 0, Max: 0},
		size:                commondata.Size{Width: 0, Height: 0},
		radiusOfDestruction: 0,
		rotation:            0,
		ricochet:            0,
		triggerRadius:       0,
		triggerDelay:        0,
		lifeTime:            0,
		isRemoved:           false,
		effects:             []effect.Effect{},
	}

	return b
}

func (b *Bullet) GetKind() string {
	return b.kind
}

func (b *Bullet) GetNickname() string {
	return b.nickname
}

func (b *Bullet) GetHealth() int {
	return b.health
}

func (b *Bullet) GetPosition() commondata.Position {
	return b.position
}

func (b *Bullet) GetTargetPosition() commondata.Position {
	return b.targetPosition
}

func (b *Bullet) GetSize() commondata.Size {
	return b.size
}

func (b *Bullet) GetSpeed() commondata.Speed {
	return b.speed
}

func (b *Bullet) GetRotation() int {
	return b.rotation
}

func (b *Bullet) GetLifeTime() int {
	return b.lifeTime
}

func (b *Bullet) GetEffects() []effect.Effect {
	return b.effects
}

func (b *Bullet) GetRadiusOfDestruction() int {
	return b.radiusOfDestruction
}

func (b *Bullet) GetTriggerRadius() int {
	return b.triggerRadius
}

func (b *Bullet) GetTriggerDelay() int {
	return b.triggerDelay
}

func (b *Bullet) GetRicochet() int {
	return b.ricochet
}

func (b *Bullet) SetKind(kind string) {
	b.kind = kind
}

func (b *Bullet) SetNickname(nickname string) {
	b.nickname = nickname
}

func (b *Bullet) SetPosition(position commondata.Position) {
	b.position = position
}

func (b *Bullet) SetTargetPosition(targetPosition commondata.Position) {
	b.targetPosition = targetPosition
}

func (b *Bullet) SetSize(size commondata.Size) {
	b.size = size
}

func (b *Bullet) SetLifeTime(lifeTime int) {
	b.lifeTime = lifeTime
}

func (b *Bullet) SetSpeed(speed commondata.Speed) {
	b.speed = speed
}

func (b *Bullet) SetRotation(rotation int) {
	b.rotation = rotation
}

func (b *Bullet) SetHealth(health int) {
	b.health = health
}

func (b *Bullet) SetEffects(effects []effect.Effect) {
	b.effects = effects
}

func (b *Bullet) SetRadiusOfDestruction(radiusOfDestruction int) {
	b.radiusOfDestruction = radiusOfDestruction
}

func (b *Bullet) SetTriggerRadius(triggerRadius int) {
	b.triggerRadius = triggerRadius
}

func (b *Bullet) SetTriggerDelay(triggerDelay int) {
	b.triggerDelay = triggerDelay
}

func (b *Bullet) SetRicochet(ricochet int) {
	b.ricochet = ricochet
}
