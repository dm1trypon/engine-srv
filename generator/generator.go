package generator

import (
	"math/rand"
	"time"

	"github.com/dm1trypon/engine-srv/config"
	"github.com/dm1trypon/engine-srv/generator/sceneobject"
	"github.com/dm1trypon/engine-srv/goapi/models/commondata"
)

type Generator struct {
	sizeCell   int
	wCells     int
	hCells     int
	objectsMap [][]*sceneobject.SceneObject
}

func (g *Generator) Create() *Generator {
	scnObj := config.ConfigInst.GetConfigData().GameObjects.Scene

	g = &Generator{
		sizeCell:   scnObj.SizeCell,
		wCells:     scnObj.WidthCountCell,
		hCells:     scnObj.HeightCountCell,
		objectsMap: [][]*sceneobject.SceneObject{},
	}

	g.makeEmptyMap()
	g.generateMap()

	return g
}

func (g *Generator) makeEmptyMap() {
	g.objectsMap = make([][]*sceneobject.SceneObject, g.wCells)
	for index := range g.objectsMap {
		g.objectsMap[index] = make([]*sceneobject.SceneObject, g.hCells)
	}
}

func (g *Generator) GetObjectsMap() [][]*sceneobject.SceneObject {
	return g.objectsMap
}

func (g *Generator) GeneratePosition() commondata.Position {
	wCell := rand.Intn(g.wCells)
	hCell := rand.Intn(g.hCells)

	for {
		if g.objectsMap[wCell][hCell] == nil {
			break
		}
	}

	return commondata.Position{
		X: wCell*g.sizeCell + g.sizeCell/2,
		Y: hCell*g.sizeCell + g.sizeCell/2,
	}
}

func (g *Generator) generateMap() {
	rand.Seed(time.Now().UTC().UnixNano())
	objLimits := config.ConfigInst.GetConfigData().GameObjects.Scene.ObjectLimits

	for _, objLimit := range objLimits {
		for count := 0; count < objLimit.Count; count++ {
			wCell := rand.Intn(g.wCells)
			hCell := rand.Intn(g.hCells)

			if g.objectsMap[wCell][hCell] != nil {
				count--
				continue
			}

			scnObj := new(sceneobject.SceneObject).Create()
			scnObj.SetMaxCount(objLimit.Count)
			scnObj.SetRespDelay(objLimit.RespDelay)
			scnObj.SetKind(objLimit.Kind)

			g.objectsMap[wCell][hCell] = scnObj
		}
	}
}
