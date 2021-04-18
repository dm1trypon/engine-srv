package module

import (
	"github.com/dm1trypon/engine-srv/goapi/models/commondata"
	"github.com/dm1trypon/engine-srv/goapi/player/weapon/module/effect"
)

func (m *Module) Create() *Module {
	m = &Module{
		kind:        "",
		lifeTime:    0,
		bulletSpeed: 0,
		reload:      0,
		damage:      0,
		bulletSize: commondata.Size{
			Width:  0,
			Height: 0,
		},
		cartridgeCapacity: 0,
		rateOfFire:        0,
		isExist:           false,
		effects:           []effect.Effect{},
	}

	return m
}

func (m *Module) GetKind() string {
	return m.kind
}

func (m *Module) GetLifeTime() int {
	return m.lifeTime
}

func (m *Module) GetBulletSpeed() int {
	return m.bulletSpeed
}

func (m *Module) GetReload() int {
	return m.reload
}

func (m *Module) GetDamage() int {
	return m.damage
}

func (m *Module) GetBulletSize() commondata.Size {
	return m.bulletSize
}

func (m *Module) GetCartridgeCapacity() int {
	return m.cartridgeCapacity
}

func (m *Module) GetRateOfFire() int {
	return m.rateOfFire
}

func (m *Module) GetIsExist() bool {
	return m.isExist
}

func (m *Module) GetEffects() []effect.Effect {
	return m.effects
}

func (m *Module) SetKind(kind string) {
	m.kind = kind
}

func (m *Module) SetLifeTime(lifeTime int) {
	m.lifeTime = lifeTime
}

func (m *Module) SetBulletSpeed(bulletSpeed int) {
	m.bulletSpeed = bulletSpeed
}

func (m *Module) SetReload(reload int) {
	m.reload = reload
}

func (m *Module) SetDamage(damage int) {
	m.damage = damage
}

func (m *Module) SetBulletSize(bulletSize commondata.Size) {
	m.bulletSize = bulletSize
}

func (m *Module) SetCartridgeCapacity(cartridgeCapacity int) {
	m.cartridgeCapacity = cartridgeCapacity
}

func (m *Module) SetRateOfFire(rateOfFire int) {
	m.rateOfFire = rateOfFire
}

func (m *Module) SetIsExist(isExist bool) {
	m.isExist = isExist
}

func (m *Module) SetEffects(effects []effect.Effect) {
	m.effects = effects
}
