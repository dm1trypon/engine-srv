package weapon

import (
	"github.com/dm1trypon/engine-srv/goapi/player/weapon/cartridge"
	"github.com/dm1trypon/engine-srv/goapi/player/weapon/module"
)

func (w *Weapon) Create() *Weapon {
	w = &Weapon{
		kind:               "",
		curRateOfFire:      0,
		maxRateOfFire:      0,
		curParts:           0,
		maxParts:           0,
		curReloadTime:      0,
		maxReloadTime:      0,
		curCountCartridges: 0,
		cartridgeCapacity:  0,
		curCartridgeStores: 0,
		maxCartridgeStores: 0,
		spread:             0,
		isActive:           false,
		isDefault:          false,
		cartridge:          cartridge.Cartridge{},
		modules:            []module.Module{},
	}

	return w
}

func (w *Weapon) GetCountCartridges() int {
	return w.curCountCartridges
}

func (w *Weapon) GetKind() string {
	return w.kind
}

func (w *Weapon) GetCurRateOfFire() int {
	return w.curRateOfFire
}

func (w *Weapon) GetMaxRateOfFire() int {
	return w.maxRateOfFire
}

func (w *Weapon) GetCurParts() int {
	return w.curParts
}

func (w *Weapon) GetMaxParts() int {
	return w.maxParts
}

func (w *Weapon) GetCurReloadTime() int {
	return w.curReloadTime
}

func (w *Weapon) GetMaxReloadTime() int {
	return w.maxReloadTime
}

func (w *Weapon) GetCurCountCartridges() int {
	return w.curCountCartridges
}

func (w *Weapon) GetCartridgeCapacity() int {
	return w.cartridgeCapacity
}

func (w *Weapon) GetCurCartridgeStores() int {
	return w.curCartridgeStores
}

func (w *Weapon) GetMaxCartridgeStores() int {
	return w.maxCartridgeStores
}

func (w *Weapon) GetSpread() int {
	return w.spread
}

func (w *Weapon) GetIsActive() bool {
	return w.isActive
}

func (w *Weapon) GetIsDefault() bool {
	return w.isDefault
}

func (w *Weapon) GetCartridge() cartridge.Cartridge {
	return w.cartridge
}

func (w *Weapon) GetModules() []module.Module {
	return w.modules
}

func (w *Weapon) SetCountCartridges(curCountCartridges int) {
	w.curCountCartridges = curCountCartridges
}

func (w *Weapon) SetKind(kind string) {
	w.kind = kind
}

func (w *Weapon) SetCurRateOfFire(curRateOfFire int) {
	w.curRateOfFire = curRateOfFire
}

func (w *Weapon) SetMaxRateOfFire(maxRateOfFire int) {
	w.maxRateOfFire = maxRateOfFire
}

func (w *Weapon) SetCurParts(curParts int) {
	w.curParts = curParts
}

func (w *Weapon) SetMaxParts(maxParts int) {
	w.maxParts = maxParts
}

func (w *Weapon) SetCurReloadTime(curReloadTime int) {
	w.curReloadTime = curReloadTime
}

func (w *Weapon) SetMaxReloadTime(maxReloadTime int) {
	w.maxReloadTime = maxReloadTime
}

func (w *Weapon) SetCurCountCartridges(curCountCartridges int) {
	w.curCountCartridges = curCountCartridges
}

func (w *Weapon) SetCartridgeCapacity(cartridgeCapacity int) {
	w.cartridgeCapacity = cartridgeCapacity
}

func (w *Weapon) SetCurCartridgeStores(curCartridgeStores int) {
	w.curCartridgeStores = curCartridgeStores
}

func (w *Weapon) SetMaxCartridgeStores(maxCartridgeStores int) {
	w.maxCartridgeStores = maxCartridgeStores
}

func (w *Weapon) SetSpread(spread int) {
	w.spread = spread
}

func (w *Weapon) SetIsActive(isActive bool) {
	w.isActive = isActive
}

func (w *Weapon) SetIsDefault(isDefault bool) {
	w.isDefault = isDefault
}

func (w *Weapon) SetCartridges(cartridge cartridge.Cartridge) {
	w.cartridge = cartridge
}

func (w *Weapon) SetModules(modules []module.Module) {
	w.modules = modules
}
