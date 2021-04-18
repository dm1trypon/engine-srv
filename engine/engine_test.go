package engine

import (
	"fmt"
	"testing"

	"github.com/dm1trypon/engine-srv/config"
	"github.com/dm1trypon/engine-srv/goapi/models/commondata"
	"github.com/dm1trypon/engine-srv/goapi/player/cursor"
)

func TestNewPlayer(t *testing.T) {
	config.ConfigInst = new(config.Config).Create()

	e := new(Engine).Create()
	e.NewPlayer("test")

	armor := e.gameObjectsData.Players[0].GetArmor()
	health := e.gameObjectsData.Players[0].GetHealth()
	nickname := e.gameObjectsData.Players[0].GetNickname()

	if nickname != "test" {
		t.Error(fmt.Sprint("Invalid parameter 'Nickname' ['test' -> ", nickname, "]"))
	}

	if armor != 200 {
		t.Error(fmt.Sprint("Invalid parameter 'Armor' [200 -> ", armor, "]"))
	}

	if health != 200 {
		t.Error(fmt.Sprint("Invalid parameter 'Health' [200 -> ", health, "]"))
	}
}

func TestBullet(t *testing.T) {
	config.ConfigInst = new(config.Config).Create()

	e := new(Engine).Create()
	e.NewPlayer("test")

	crsrObj := new(cursor.Cursor).Create()
	crsrObj.SetIsPressed(true)

	e.gameObjectsData.Players[0].SetCursor(*crsrObj)
	e.gameObjectsData.Players[0].SetPosition(commondata.Position{
		X: 100,
		Y: 100,
	})

	e.gameObjectsData.Players[0].SetRotation(45)

	wpnsList := e.gameObjectsData.Players[0].GetWeapons()

	for index := 0; index < len(wpnsList); index++ {
		if wpnsList[index].GetIsDefault() {
			wpnsList[index].SetIsActive(true)
		}
	}

	e.gameObjectsData.Players[0].SetWeapons(wpnsList)

	e.NewBullet(&e.gameObjectsData.Players[0])
	blltTrgtPos := e.gameObjectsData.Bullets[0].GetTargetPosition()
	blltSpeed := e.gameObjectsData.Bullets[0].GetSpeed()
	blltRotation := e.gameObjectsData.Bullets[0].GetRotation()

	if blltTrgtPos.X != 630 || blltTrgtPos.Y != 630 {
		t.Error(fmt.Sprint("Invalid parameter 'TargetPosition' [-650, 100] -> [",
			blltTrgtPos.X, ", ", blltTrgtPos.Y, "]"))
	}

	if blltSpeed.X != 28 || blltSpeed.Y != 28 || blltSpeed.Max != 40 {
		t.Error(fmt.Sprint("Invalid parameter 'Speed' [28, 28, 40] -> [",
			blltSpeed.X, ", ", blltSpeed.Y, ", ", blltSpeed.Max, "]"))
	}

	if blltRotation != 45 {
		t.Error(fmt.Sprint("Invalid parameter 'Rotation' [45] -> [", blltRotation, "]"))
	}
}

func TestOnPlayerBulletsCollision(t *testing.T) {
	config.ConfigInst = new(config.Config).Create()

	e := new(Engine).Create()
	e.NewPlayer("test")

	crsrObj := new(cursor.Cursor).Create()
	crsrObj.SetIsPressed(true)

	e.gameObjectsData.Players[0].SetCursor(*crsrObj)
	e.gameObjectsData.Players[0].SetPosition(commondata.Position{
		X: 0,
		Y: 0,
	})

	e.gameObjectsData.Players[0].SetRotation(45)

	wpnsList := e.gameObjectsData.Players[0].GetWeapons()

	for index := 0; index < len(wpnsList); index++ {
		if wpnsList[index].GetIsDefault() {
			wpnsList[index].SetIsActive(true)
		}
	}

	e.gameObjectsData.Players[0].SetWeapons(wpnsList)

	e.NewBullet(&e.gameObjectsData.Players[0])
	e.gameObjectsData.Bullets[0].SetPosition(commondata.Position{X: 0, Y: -111})

	e.onPlayerBulletsCollision(&e.gameObjectsData.Players[0])
}
