package engine

import (
	"fmt"
	"math"
	"time"

	logger "github.com/dm1trypon/easy-logger"
	"github.com/dm1trypon/engine-srv/collision"
	"github.com/dm1trypon/engine-srv/config"
	"github.com/dm1trypon/engine-srv/generator"
	"github.com/dm1trypon/engine-srv/goapi/ammocase"
	"github.com/dm1trypon/engine-srv/goapi/animation"
	"github.com/dm1trypon/engine-srv/goapi/block"
	"github.com/dm1trypon/engine-srv/goapi/bullet"
	"github.com/dm1trypon/engine-srv/goapi/guncase"
	"github.com/dm1trypon/engine-srv/goapi/gunchest"
	"github.com/dm1trypon/engine-srv/goapi/gunmodule"
	"github.com/dm1trypon/engine-srv/goapi/gunpart"
	"github.com/dm1trypon/engine-srv/goapi/loot"
	"github.com/dm1trypon/engine-srv/goapi/models/commondata"
	"github.com/dm1trypon/engine-srv/goapi/player"
	plEffect "github.com/dm1trypon/engine-srv/goapi/player/effect"
	"github.com/dm1trypon/engine-srv/goapi/scene"
	"github.com/dm1trypon/engine-srv/goapi/specialrune"
	"github.com/dm1trypon/engine-srv/randomizer"
)

// LC - logging's category
const LC = "ENGINE"

type Engine struct {
	randomizerInst  *randomizer.Randomizer
	generatorInst   *generator.Generator
	gameObjectsData goData
	maxTimeFPS      int
	maxTimeStep     int
	msFPS           int
	msTimeStep      int
}

type goData struct {
	Players      []player.Player
	Bullets      []bullet.Bullet
	Blocks       []block.Block
	AmmoCases    []ammocase.AmmoCase
	Animations   []animation.Animation
	GunCases     []guncase.GunCase
	GunChests    []gunchest.GunChest
	GunModules   []gunmodule.GunModule
	GunParts     []gunpart.GunPart
	Loots        []loot.Loot
	Scenes       []scene.Scene
	SpecialRunes []specialrune.SpecialRune
}

func (e *Engine) Create() *Engine {
	logger.Info(LC, "Creating Engine module")

	e.randomizerInst = new(randomizer.Randomizer).Create()
	e.generatorInst = new(generator.Generator).Create()

	engineCfg := config.ConfigInst.GetConfigData().Engine
	e.maxTimeFPS = engineCfg.FPS
	e.maxTimeStep = engineCfg.TimeStep
	e.msFPS = 0
	e.msTimeStep = 0

	e.initGameObjects()
	go e.startLoop()

	return e
}

func (e *Engine) initGameObjects() {
	e.gameObjectsData = goData{
		Players:      []player.Player{},
		Bullets:      []bullet.Bullet{},
		Blocks:       []block.Block{},
		AmmoCases:    []ammocase.AmmoCase{},
		Animations:   []animation.Animation{},
		GunCases:     []guncase.GunCase{},
		GunChests:    []gunchest.GunChest{},
		GunModules:   []gunmodule.GunModule{},
		GunParts:     []gunpart.GunPart{},
		Loots:        []loot.Loot{},
		Scenes:       []scene.Scene{},
		SpecialRunes: []specialrune.SpecialRune{},
	}
}

func (e *Engine) GetGameObjectsData() *goData {
	return &e.gameObjectsData
}

func (e *Engine) calcBulletSpeed(blltObj *bullet.Bullet) {
	blltPos := blltObj.GetPosition()
	blltTrgtPos := blltObj.GetTargetPosition()
	blltSpeed := blltObj.GetSpeed()

	fmt.Println(blltPos, blltTrgtPos, blltSpeed)

	speedX := int((blltTrgtPos.X - blltPos.X) * blltSpeed.Max /
		int(math.Sqrt(math.Abs(float64((blltTrgtPos.X-blltPos.X)*(blltTrgtPos.X-blltPos.X)+
			(blltTrgtPos.Y-blltPos.Y)*(blltTrgtPos.Y-blltPos.Y))))))

	speedY := int((blltTrgtPos.Y - blltPos.Y) * blltSpeed.Max /
		int(math.Sqrt(math.Abs(float64((blltTrgtPos.X-blltPos.X)*(blltTrgtPos.X-blltPos.X)+
			(blltTrgtPos.Y-blltPos.Y)*(blltTrgtPos.Y-blltPos.Y))))))

	blltSpeed.X = speedX
	blltSpeed.Y = speedY

	blltObj.SetSpeed(blltSpeed)
}

func (e *Engine) calcBulletTargetPosition(blltObj *bullet.Bullet, plObj *player.Player, spread int) {
	plPos := plObj.GetPosition()
	plRotation := plObj.GetRotation()

	blltSpeed := blltObj.GetSpeed().Max
	blltLifeTime := blltObj.GetLifeTime()

	rad := float64(plRotation+spread/2) * math.Pi / 180

	blltDist := blltSpeed * blltLifeTime
	distPerInt := blltDist / 8

	trgtPosInt := commondata.Position{
		X: plPos.X + int(float64(distPerInt)*math.Cos(rad)),
		Y: plPos.Y + int(float64(distPerInt)*math.Sin(rad)),
	}

	blltObj.SetTargetPosition(trgtPosInt)
}

func (e *Engine) reincarnationPlayer(nickname string) {

}

func (e *Engine) startLoop() {
	for {
		e.nextFrame()
		e.msFPS++
		e.msTimeStep++
		time.Sleep(time.Millisecond)
	}
}

func (e *Engine) nextFrame() {
	if e.msFPS == e.maxTimeFPS {
		e.onFPSStep()
		e.msFPS = 0
	}

	if e.msTimeStep == e.maxTimeStep {
		e.onTimeStep()
		e.msTimeStep = 0
	}
}

func (e *Engine) onFPSStep() {
	e.collisionChecking()
	// calculationPlayers()
	// calculationBullets()
}

func (e *Engine) collisionChecking() {
	for iPlObj := 0; iPlObj < len(e.gameObjectsData.Players); iPlObj++ {
		plObjs := &e.gameObjectsData.Players[iPlObj]

		e.onPlayerBulletsCollision(plObjs)
		e.onPlayerAmmoCaseCollision(plObjs)
		e.onPlayerLootsCollision(plObjs)
		e.onPlayerRunesCollision(plObjs)
		e.onPlayerSceneCollision(plObjs)
		e.onPlayerBlocksCollision(plObjs)
		e.onPlayerGunCasesCollision(plObjs)
		e.onPlayerGunPartCollision(plObjs)
		e.onPlayerGunModuleCollision(plObjs)
		e.onPlayerGunChestCollision(plObjs)
	}
}

func (e *Engine) onPlayerBulletsCollision(plObj *player.Player) {
	for index := 0; index < len(e.gameObjectsData.Bullets); index++ {
		blltObj := &e.gameObjectsData.Bullets[index]

		plPosition := plObj.GetPosition()
		plSize := plObj.GetSize()

		blltPosition := blltObj.GetPosition()
		blltSize := blltObj.GetSize()

		pair := collision.CircCircPair{
			CircleOne: collision.Circle{
				Radius:   float64(plSize.Width) / 2,
				Position: plPosition,
			},
			CircleTwo: collision.Circle{
				Radius:   float64(blltSize.Height) / 2,
				Position: blltPosition,
			},
		}

		colInst := new(collision.Collision).Create(pair)

		if colInst.Check() < 0 {
			continue
		}

		logger.Debug(LC, "Found collision Player with Bullet")

		plHealth := plObj.GetHealth()
		blltHealth := blltObj.GetHealth()

		plObj.SetHealth(plHealth - blltHealth)
		blltObj.SetHealth(blltHealth - plHealth)
	}
}

func (e *Engine) onPlayerSceneCollision(plObj *player.Player) {
	for _, scnObj := range e.gameObjectsData.Scenes {
		plPosition := plObj.GetPosition()
		plSize := plObj.GetSize()

		scnPosition := scnObj.GetPosition()
		scnSize := scnObj.GetSize()

		pair := collision.CircRectPair{
			Circle: collision.Circle{
				Radius:   float64(plSize.Width) / 2,
				Position: plPosition,
			},
			Rectangle: collision.Rectangle{
				Size:     scnSize,
				Position: scnPosition,
				Segments: []collision.Segment{},
			},
		}

		colInst := new(collision.Collision).Create(pair)
		sectCol := colInst.Check()

		if sectCol < 0 {
			continue
		}

		logger.Debug(LC, "Found collision Player with Scene")

		e.bound(plObj, sectCol)
	}
}

func (e *Engine) bound(plObj *player.Player, sectCol int) {
	plPos := plObj.GetPosition()
	plSpeed := plObj.GetSpeed()

	if sectCol == 0 {
		plSpeed.X = 0
		plPos.X = plPos.X - plSpeed.X

		plObj.SetSpeed(plSpeed)
		plObj.SetPosition(plPos)
	} else if sectCol == 1 {
		plSpeed.Y = 0
		plPos.Y = plPos.Y - plSpeed.Y

		plObj.SetSpeed(plSpeed)
		plObj.SetPosition(plPos)
	} else if sectCol == 2 {
		plSpeed.X = 0
		plPos.X = plPos.X + plSpeed.X

		plObj.SetSpeed(plSpeed)
		plObj.SetPosition(plPos)
	} else if sectCol == 3 {
		plSpeed.Y = 0
		plPos.Y = plPos.Y + plSpeed.Y

		plObj.SetSpeed(plSpeed)
		plObj.SetPosition(plPos)
	}
}

func (e *Engine) onPlayerBlocksCollision(plObj *player.Player) {
	for _, blckObj := range e.gameObjectsData.Blocks {
		plPosition := plObj.GetPosition()
		plSize := plObj.GetSize()

		blckPosition := blckObj.GetPosition()
		blckSize := blckObj.GetSize()

		pair := collision.CircRectPair{
			Circle: collision.Circle{
				Radius:   float64(plSize.Width) / 2,
				Position: plPosition,
			},
			Rectangle: collision.Rectangle{
				Size:     blckSize,
				Position: blckPosition,
				Segments: []collision.Segment{},
			},
		}

		colInst := new(collision.Collision).Create(pair)
		sectCol := colInst.Check()

		if sectCol < 0 {
			continue
		}

		e.bound(plObj, sectCol)
	}
}

func (e *Engine) onPlayerLootsCollision(plObj *player.Player) {
	for index := 0; index < len(e.gameObjectsData.Loots); index++ {
		ltObj := &e.gameObjectsData.Loots[index]

		plPosition := plObj.GetPosition()
		plSize := plObj.GetSize()

		ltPosition := ltObj.GetPosition()
		ltSize := ltObj.GetSize()

		pair := collision.CircRectPair{
			Circle: collision.Circle{
				Radius:   float64(plSize.Width) / 2,
				Position: plPosition,
			},
			Rectangle: collision.Rectangle{
				Size:     ltSize,
				Position: ltPosition,
				Segments: []collision.Segment{},
			},
		}

		colInst := new(collision.Collision).Create(pair)
		sectCol := colInst.Check()

		if sectCol < 0 {
			continue
		}

		ltKind := ltObj.GetKind()
		ltValue := ltObj.GetValue()

		if ltKind == "first_aid_kit" {
			plHealth := plObj.GetHealth()
			plObj.SetHealth(plHealth + ltValue)
		} else if ltKind == "armor" {
			plArmor := plObj.GetArmor()
			plObj.SetArmor(plArmor + ltValue)
		} else {
			plWeapons := plObj.GetWeapons()

			for _, wpnObj := range plWeapons {
				wpnKind := wpnObj.GetKind()

				if wpnKind != ltKind {
					continue
				}

				wpnCurCountCartridges := wpnObj.GetCountCartridges()
				cartriges := ltValue + wpnCurCountCartridges

				for {
					wpnCartridgeCapacity := wpnObj.GetCartridgeCapacity()

					if cartriges > wpnCartridgeCapacity {
						wpnCartStrores := wpnObj.GetCurCartridgeStores() + 1
						wpnObj.SetCurCartridgeStores(wpnCartStrores)
						wpnCartCap := wpnObj.GetCartridgeCapacity()
						wpnObj.SetCountCartridges(wpnCartCap)
						cartriges -= wpnCartCap
					} else {
						wpnObj.SetCountCartridges(cartriges)
						break
					}
				}
			}

			plObj.SetWeapons(plWeapons)
		}

		e.removeLoot(index)
	}
}

func (e *Engine) onPlayerAmmoCaseCollision(plObj *player.Player) {
	for index := 0; index < len(e.gameObjectsData.AmmoCases); index++ {
		amCsObj := &e.gameObjectsData.AmmoCases[index]

		plPos := plObj.GetPosition()
		plSize := plObj.GetSize()

		amCsPos := amCsObj.GetPosition()
		amCsSize := amCsObj.GetSize()

		pair := collision.CircRectPair{
			Circle: collision.Circle{
				Radius:   float64(plSize.Width) / 2,
				Position: plPos,
			},
			Rectangle: collision.Rectangle{
				Size:     amCsSize,
				Position: amCsPos,
				Segments: []collision.Segment{},
			},
		}

		colInst := new(collision.Collision).Create(pair)
		sectCol := colInst.Check()

		if sectCol < 0 {
			continue
		}

		e.bound(plObj, sectCol)
	}
}

func (e *Engine) onPlayerRunesCollision(plObj *player.Player) {
	for index := 0; index < len(e.gameObjectsData.SpecialRunes); index++ {
		spclRnObj := &e.gameObjectsData.SpecialRunes[index]

		plPos := plObj.GetPosition()
		plSize := plObj.GetSize()

		spclRnPos := spclRnObj.GetPosition()
		spclRnSize := spclRnObj.GetSize()

		pair := collision.CircRectPair{
			Circle: collision.Circle{
				Radius:   float64(plSize.Width) / 2,
				Position: plPos,
			},
			Rectangle: collision.Rectangle{
				Size:     spclRnSize,
				Position: spclRnPos,
				Segments: []collision.Segment{},
			},
		}

		colInst := new(collision.Collision).Create(pair)
		sectCol := colInst.Check()

		if sectCol < 0 {
			continue
		}

		spclRnEfObj := spclRnObj.GetEffect()
		plEffects := plObj.GetEffects()

		spclRnEfKind := spclRnEfObj.GetKind()
		spclRnEfDur := spclRnEfObj.GetDuration()
		spclRnEfInvis := spclRnEfObj.GetIsInvisibility()
		spclRnEfSpeed := spclRnEfObj.GetSpeed()

		for _, plEfObj := range plEffects {
			plEfKind := plEfObj.GetKind()

			if plEfKind != spclRnEfKind {
				continue
			}

			plEfObj.SetIsInvisibility(spclRnEfInvis)
			plEfObj.SetSpeed(spclRnEfSpeed)
			plEfObj.SetDuration(spclRnEfDur)
			plEfObj.SetKind(spclRnEfKind)

			e.removeSpecialRune(index)
			return
		}

		plEfObj := new(plEffect.Effect).Create()
		plEfObj.SetIsInvisibility(spclRnEfInvis)
		plEfObj.SetSpeed(spclRnEfSpeed)
		plEfObj.SetDuration(spclRnEfDur)
		plEfObj.SetKind(spclRnEfKind)

		plEffects = append(plEffects, *plEfObj)
		plObj.SetEffects(plEffects)

		e.removeSpecialRune(index)
	}
}

func (e *Engine) onPlayerGunCasesCollision(plObj *player.Player) {
	for index := 0; index < len(e.gameObjectsData.GunCases); index++ {
		gnCsObj := &e.gameObjectsData.GunCases[index]

		plPos := plObj.GetPosition()
		plSize := plObj.GetSize()

		gnCsPos := gnCsObj.GetPosition()
		gnCsSize := gnCsObj.GetSize()

		pair := collision.CircRectPair{
			Circle: collision.Circle{
				Radius:   float64(plSize.Width) / 2,
				Position: plPos,
			},
			Rectangle: collision.Rectangle{
				Size:     gnCsSize,
				Position: gnCsPos,
				Segments: []collision.Segment{},
			},
		}

		colInst := new(collision.Collision).Create(pair)
		sectCol := colInst.Check()

		if sectCol < 0 {
			continue
		}

		e.bound(plObj, sectCol)
	}
}

func (e *Engine) onPlayerGunPartCollision(plObj *player.Player) {
	for index := 0; index < len(e.gameObjectsData.GunParts); index++ {
		gnPrtObj := &e.gameObjectsData.GunParts[index]

		plPos := plObj.GetPosition()
		plSize := plObj.GetSize()

		gnPrtPos := gnPrtObj.GetPosition()
		gnPrtSize := gnPrtObj.GetSize()

		pair := collision.CircRectPair{
			Circle: collision.Circle{
				Radius:   float64(plSize.Width) / 2,
				Position: plPos,
			},
			Rectangle: collision.Rectangle{
				Size:     gnPrtSize,
				Position: gnPrtPos,
				Segments: []collision.Segment{},
			},
		}

		colInst := new(collision.Collision).Create(pair)
		sectCol := colInst.Check()

		if sectCol < 0 {
			continue
		}

		plWpnList := plObj.GetWeapons()
		gnPrtKind := gnPrtObj.GetKind()

		for _, wpnObj := range plWpnList {
			wpnKind := wpnObj.GetKind()

			if wpnKind != gnPrtKind {
				continue
			}

			wpnCurParts := wpnObj.GetCurParts()
			wpnMaxParts := wpnObj.GetMaxParts()

			if wpnCurParts >= wpnMaxParts {
				return
			}

			wpnObj.SetCurParts(wpnCurParts + 1)
		}

		plObj.SetWeapons(plWpnList)

		e.removeGunPart(index)
	}
}

func (e *Engine) onPlayerGunModuleCollision(plObj *player.Player) {
	for index := 0; index < len(e.gameObjectsData.GunModules); index++ {
		gnMdlObj := &e.gameObjectsData.GunModules[index]

		plPos := plObj.GetPosition()
		plSize := plObj.GetSize()

		gnMdlPos := gnMdlObj.GetPosition()
		gnMdlSize := gnMdlObj.GetSize()

		pair := collision.CircRectPair{
			Circle: collision.Circle{
				Radius:   float64(plSize.Width) / 2,
				Position: plPos,
			},
			Rectangle: collision.Rectangle{
				Size:     gnMdlSize,
				Position: gnMdlPos,
				Segments: []collision.Segment{},
			},
		}

		colInst := new(collision.Collision).Create(pair)
		sectCol := colInst.Check()

		if sectCol < 0 {
			continue
		}

		plWpnList := plObj.GetWeapons()
		gnPrtKind := gnMdlObj.GetKind()

		for _, wpnObj := range plWpnList {
			wpnMdlList := wpnObj.GetModules()

			for _, mdlObj := range wpnMdlList {
				mdlKind := mdlObj.GetKind()

				if mdlKind != gnPrtKind {
					continue
				}

				mdlIsExist := mdlObj.GetIsExist()

				if mdlIsExist {
					continue
				}

				mdlObj.SetIsExist(true)
			}

			wpnObj.SetModules(wpnMdlList)
		}

		plObj.SetWeapons(plWpnList)

		e.removeGunModule(index)
	}
}

func (e *Engine) onPlayerGunChestCollision(plObj *player.Player) {
	for index := 0; index < len(e.gameObjectsData.GunChests); index++ {
		gnChstObj := &e.gameObjectsData.GunChests[index]

		plPos := plObj.GetPosition()
		plSize := plObj.GetSize()

		gnChstPos := gnChstObj.GetPosition()
		gnChstSize := gnChstObj.GetSize()

		pair := collision.CircRectPair{
			Circle: collision.Circle{
				Radius:   float64(plSize.Width) / 2,
				Position: plPos,
			},
			Rectangle: collision.Rectangle{
				Size:     gnChstSize,
				Position: gnChstPos,
				Segments: []collision.Segment{},
			},
		}

		colInst := new(collision.Collision).Create(pair)
		sectCol := colInst.Check()

		if sectCol < 0 {
			continue
		}

		e.bound(plObj, sectCol)
	}
}

func (e *Engine) onTimeStep() {
}

func (e *Engine) removeLoot(index int) {
	e.gameObjectsData.Loots[len(e.gameObjectsData.Loots)-1],
		e.gameObjectsData.Loots[index] = e.gameObjectsData.Loots[index],
		e.gameObjectsData.Loots[len(e.gameObjectsData.Loots)-1]

	e.gameObjectsData.Loots = e.gameObjectsData.Loots[:len(e.gameObjectsData.Loots)-1]
}

func (e *Engine) removeSpecialRune(index int) {
	e.gameObjectsData.SpecialRunes[len(e.gameObjectsData.SpecialRunes)-1],
		e.gameObjectsData.SpecialRunes[index] = e.gameObjectsData.SpecialRunes[index],
		e.gameObjectsData.SpecialRunes[len(e.gameObjectsData.SpecialRunes)-1]

	e.gameObjectsData.SpecialRunes = e.gameObjectsData.SpecialRunes[:len(e.gameObjectsData.SpecialRunes)-1]
}

func (e *Engine) removeGunPart(index int) {
	e.gameObjectsData.GunParts[len(e.gameObjectsData.GunParts)-1],
		e.gameObjectsData.GunParts[index] = e.gameObjectsData.GunParts[index],
		e.gameObjectsData.GunParts[len(e.gameObjectsData.GunParts)-1]

	e.gameObjectsData.GunParts = e.gameObjectsData.GunParts[:len(e.gameObjectsData.GunParts)-1]
}

func (e *Engine) removeGunModule(index int) {
	e.gameObjectsData.GunModules[len(e.gameObjectsData.GunModules)-1],
		e.gameObjectsData.GunModules[index] = e.gameObjectsData.GunModules[index],
		e.gameObjectsData.GunModules[len(e.gameObjectsData.GunModules)-1]

	e.gameObjectsData.GunModules = e.gameObjectsData.GunModules[:len(e.gameObjectsData.GunModules)-1]
}
