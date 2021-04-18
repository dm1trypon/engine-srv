package engine

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
	logger "github.com/dm1trypon/easy-logger"
	"github.com/dm1trypon/engine-srv/config"
	"github.com/dm1trypon/engine-srv/goapi/ammocase"
	"github.com/dm1trypon/engine-srv/goapi/bullet"
	blltEffect "github.com/dm1trypon/engine-srv/goapi/bullet/effect"
	"github.com/dm1trypon/engine-srv/goapi/models/commondata"
	"github.com/dm1trypon/engine-srv/goapi/player"
	"github.com/dm1trypon/engine-srv/goapi/player/weapon"
	"github.com/dm1trypon/engine-srv/goapi/player/weapon/cartridge"
	crtrdgEffect "github.com/dm1trypon/engine-srv/goapi/player/weapon/cartridge/effect"
	"github.com/dm1trypon/engine-srv/goapi/player/weapon/module"
)

func (e *Engine) NewPlayer(nickname string) {
	goCfg := config.ConfigInst.GetConfigData().GameObjects
	plCfgObj := goCfg.Players
	wpnsCfgObj := goCfg.Weapons

	var wpnList []weapon.Weapon

	for _, wpnCfgObj := range wpnsCfgObj {
		wpnObj := new(weapon.Weapon).Create()

		wpnObj.SetSpread(wpnCfgObj.BulletSpread)
		wpnObj.SetMaxParts(wpnCfgObj.MaxParts)
		wpnObj.SetCartridgeCapacity(wpnCfgObj.CartridgeCapacity)
		wpnObj.SetMaxRateOfFire(wpnCfgObj.RateOfFire)
		wpnObj.SetMaxReloadTime(wpnCfgObj.ReloadTime)
		wpnObj.SetMaxCartridgeStores(wpnCfgObj.CartridgeStores)
		wpnObj.SetIsDefault(wpnCfgObj.IsDefault)
		wpnObj.SetKind(wpnCfgObj.Kind)

		crtrdgObj := new(cartridge.Cartridge).Create()
		crtrdgObj.SetDamage(wpnCfgObj.BulletDamage)
		crtrdgObj.SetLifeTime(wpnCfgObj.BulletLifeTime)
		crtrdgObj.SetRadiusOfDestruction(wpnCfgObj.BulletRadiusOfDestruction)
		crtrdgObj.SetRicochet(wpnCfgObj.BulletRicochet)

		var crtrdgEfList []crtrdgEffect.Effect

		for _, efCfgObj := range wpnCfgObj.BulletEffects {
			efObj := new(crtrdgEffect.Effect).Create()
			efObj.SetChance(efCfgObj.Chance)
			efObj.SetDuration(efCfgObj.Duration)
			efObj.SetKind(efCfgObj.Kind)
			efObj.SetRateOfFire(efCfgObj.Rate)
			efObj.SetSpeed(efCfgObj.Speed)

			crtrdgEfList = append(crtrdgEfList, *efObj)
		}

		crtrdgObj.SetEffects(crtrdgEfList)

		blltSize := commondata.Size{
			Width:  wpnCfgObj.BulletWidth,
			Height: wpnCfgObj.BulletWidth,
		}

		crtrdgObj.SetSize(blltSize)
		crtrdgObj.SetTriggerDelay(wpnCfgObj.BulletTriggerDelay)
		crtrdgObj.SetTriggerRadius(wpnCfgObj.BulletTriggerRadius)
		crtrdgObj.SetSpeed(wpnCfgObj.BulletSpeed)

		wpnObj.SetCartridges(*crtrdgObj)

		var mdlList []module.Module

		for _, wpnCfgModule := range wpnCfgObj.Modules {
			mdlObj := new(module.Module).Create()

			blltSize := commondata.Size{
				Width:  wpnCfgModule.BulletWidth,
				Height: wpnCfgModule.BulletHeight,
			}

			mdlObj.SetBulletSize(blltSize)
			mdlObj.SetBulletSpeed(wpnCfgModule.BulletSpread)
			mdlObj.SetCartridgeCapacity(wpnCfgModule.CartridgeCapacity)
			mdlObj.SetDamage(wpnCfgModule.BulletDamage)
			mdlObj.SetIsExist(false)
			mdlObj.SetKind(wpnCfgModule.Kind)
			mdlObj.SetLifeTime(wpnCfgModule.BulletLifeTime)
			mdlObj.SetReload(wpnCfgModule.ReloadTime)
			mdlObj.SetRateOfFire(wpnCfgModule.RateOfFire)

			mdlList = append(mdlList, *mdlObj)
		}

		wpnObj.SetModules(mdlList)

		wpnList = append(wpnList, *wpnObj)
	}

	plPos := e.generatorInst.GeneratePosition()

	plObj := new(player.Player).Create()
	plObj.SetArmor(plCfgObj.MaxArmor)
	plObj.SetHealth(plCfgObj.MaxHealth)
	plObj.SetPosition(plPos)
	plObj.SetNickname(nickname)
	plObj.SetWeapons(wpnList)

	plSize := commondata.Size{
		Width:  plCfgObj.Width,
		Height: plCfgObj.Height,
	}

	plObj.SetSize(plSize)

	logger.Debug(LC, fmt.Sprint("[PLAYER]\n", spew.Sdump(plObj)))

	e.gameObjectsData.Players = append(e.gameObjectsData.Players, *plObj)
}

func (e *Engine) RemovePlayer(nickname string) {

}

func (e *Engine) NewBullet(plObj *player.Player) {
	plWpnList := plObj.GetWeapons()

	for _, wpnObj := range plWpnList {
		isActive := wpnObj.GetIsActive()

		if !isActive {
			continue
		}

		wpnKind := wpnObj.GetKind()
		wpnSpread := e.randomizerInst.GetRandSpread(wpnObj.GetSpread())

		crtObj := wpnObj.GetCartridge()

		plNickname := plObj.GetNickname()
		plPos := plObj.GetPosition()
		plRotate := plObj.GetRotation()

		crtDamage := crtObj.GetDamage()
		crtLifeTime := crtObj.GetLifeTime()
		crtRicochet := crtObj.GetRicochet()
		crtRadOfDistr := crtObj.GetRadiusOfDestruction()
		crtSize := crtObj.GetSize()
		crtTrigDelay := crtObj.GetTriggerDelay()
		crtTrigRad := crtObj.GetTriggerRadius()
		crtSpeed := crtObj.GetSpeed()
		crtEffects := crtObj.GetEffects()

		blltEfList := []blltEffect.Effect{}

		for _, efObj := range crtEffects {
			efChance := efObj.GetChance()

			if !e.randomizerInst.IsChance(efChance) {
				continue
			}

			blltEfObj := blltEffect.Effect{}

			blltEfObj.SetDuration(efObj.GetDuration())
			blltEfObj.SetKind(efObj.GetKind())
			blltEfObj.SetRateOfFire(efObj.GetRateOfFire())
			blltEfObj.SetSpeed(efObj.GetSpeed())

			blltEfList = append(blltEfList, blltEfObj)
		}

		blltSpeed := commondata.Speed{
			X:   0,
			Y:   0,
			Max: crtSpeed,
		}

		blltObj := new(bullet.Bullet).Create()
		blltObj.SetNickname(plNickname)
		blltObj.SetKind(wpnKind)
		blltObj.SetHealth(crtDamage)
		blltObj.SetLifeTime(crtLifeTime)
		blltObj.SetPosition(plPos)
		blltObj.SetSize(crtSize)
		blltObj.SetSpeed(blltSpeed)
		blltObj.SetRotation(plRotate)
		blltObj.SetRicochet(crtRicochet)
		blltObj.SetRadiusOfDestruction(crtRadOfDistr)
		blltObj.SetTriggerDelay(crtTrigDelay)
		blltObj.SetTriggerRadius(crtTrigRad)
		blltObj.SetEffects(blltEfList)

		e.calcBulletTargetPosition(blltObj, plObj, wpnSpread)
		e.calcBulletSpeed(blltObj)

		logger.Debug(LC, fmt.Sprint("[BULLET]\n", spew.Sdump(blltObj)))

		e.gameObjectsData.Bullets = append(e.gameObjectsData.Bullets, *blltObj)
	}
}

func (e *Engine) NewAmmoCase(kind string) {
	amCsList := config.ConfigInst.GetConfigData().GameObjects.AmmoCases

	for _, amCsCfgObj := range amCsList {
		if amCsCfgObj.Kind != kind {
			continue
		}

		amCsObj := new(ammocase.AmmoCase).Create()

		amCsSize := commondata.Size{
			Width:  amCsCfgObj.Width,
			Height: amCsCfgObj.Width,
		}

		amCsObj.SetSize(amCsSize)
		amCsObj.SetLootList(amCsCfgObj.LootList)
		amCsObj.SetHealth(amCsCfgObj.Health)
		amCsObj.SetChance(amCsCfgObj.Chance)
		amCsPos := e.generatorInst.GeneratePosition()
		amCsObj.SetPosition(amCsPos)
		amCsObj.SetKind(kind)

		e.gameObjectsData.AmmoCases = append(e.gameObjectsData.AmmoCases, *amCsObj)
	}
}
