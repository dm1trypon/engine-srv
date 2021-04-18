package physic

import (
	"math"

	"github.com/dm1trypon/engine-srv/goapi/models/commondata"
)

type Physic struct{}

func (p *Physic) BoundCircleRectangle(circPos, rectPos, directPos commondata.Position,
	distance float64) commondata.Position {

	return commondata.Position{
		X: 0,
		Y: 0,
	}
}

func (p *Physic) BoundCircles(circOnePos, circTwoPos, directPos commondata.Position,
	distance float64) commondata.Position {

	lineA := math.Sqrt(math.Pow((float64(directPos.X)-float64(circTwoPos.X)), 2) +
		+math.Pow((float64(directPos.Y)-float64(circTwoPos.Y)), 2))

	lineB := math.Sqrt(math.Pow((float64(directPos.X)-float64(circOnePos.X)), 2) +
		+math.Pow((float64(directPos.Y)-float64(circOnePos.Y)), 2))

	cosA := (math.Pow(distance, 2) + math.Pow(lineB, 2) - math.Pow(lineA, 2)) /
		(2 * distance * lineB)

	cosB := float64(directPos.X) / lineB

	if cosB > 1 {
		cosB = -1
	} else if cosB < -1 {
		cosB = 1
	}

	degE := math.Acos(cosB)*(180/math.Pi) - math.Abs((90-math.Acos(cosA)*(180/math.Pi))*2)

	posX := lineB*math.Cos(degE*(math.Pi/180)) + float64(circOnePos.X)
	posY := lineB*math.Sin(degE*(math.Pi/180)) + float64(circOnePos.Y)

	return commondata.Position{
		X: int(posX),
		Y: int(posY),
	}
}
