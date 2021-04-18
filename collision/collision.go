package collision

import (
	"math"

	"github.com/dm1trypon/engine-srv/goapi/models/commondata"
)

const LC = "COLLISION"

type Collision struct {
	pairObjects interface{}
}

func (c *Collision) Create(pairObjects interface{}) *Collision {
	c = &Collision{
		pairObjects: pairObjects,
	}

	return c
}

func (c *Collision) GetPairObjects() interface{} {
	return c.pairObjects
}

func (c *Collision) Check() int {
	switch c.pairObjects.(type) {
	case CircCircPair:
		return c.isIntersectCircleCircle()
	case CircRectPair:
		return c.isIntersectCircleRectangle()
	}

	return -1
}

func (c *Collision) isIntersectCircleCircle() int {
	pair := c.pairObjects.(CircCircPair)

	distance := math.Sqrt(math.Pow((float64(pair.CircleOne.Position.X)-float64(pair.CircleTwo.Position.X)), 2) +
		+math.Pow((float64(pair.CircleOne.Position.Y)-float64(pair.CircleTwo.Position.Y)), 2))

	if distance < float64(pair.CircleOne.Radius)+float64(pair.CircleTwo.Radius) {
		return 0
	}

	return -1
}

func (c *Collision) isIntersectCircleRectangle() int {
	pair := c.pairObjects.(CircRectPair)
	rect := pair.Rectangle
	circle := pair.Circle
	rect.makeSegments()

	for index, segment := range rect.Segments {
		if c.isIntersectCircleLine(circle, segment) {
			return index
		}
	}

	return -1
}

func (c *Collision) isIntersectCircleLine(circle Circle, segment Segment) bool {
	x01 := segment.PointOne.X - circle.Position.X
	y01 := segment.PointOne.Y - circle.Position.Y
	x02 := segment.PointTwo.X - circle.Position.X
	y02 := segment.PointTwo.Y - circle.Position.Y

	dx := x02 - x01
	dy := y02 - y01

	a := dx*dx + dy*dy
	b := 2 * (x01*dx + y01*dy)
	d := int(math.Pow(float64(x01), 2) + math.Pow(float64(y01), 2) - math.Pow(float64(circle.Radius), 2))

	if -b < 0 {
		return d < 0
	}

	if -b < (2 * a) {
		return (4*a*d - b*b) < 0
	}

	return (a + b + d) < 0
}

func (r *Rectangle) makeSegments() {
	r.Segments = []Segment{
		{
			PointOne: commondata.Position{
				X: r.Position.X - r.Size.Width/2,
				Y: r.Position.Y - r.Size.Height/2,
			},
			PointTwo: commondata.Position{
				X: r.Position.X + r.Size.Width/2,
				Y: r.Position.Y - r.Size.Height/2,
			},
		},
		{
			PointOne: commondata.Position{
				X: r.Position.X + r.Size.Width/2,
				Y: r.Position.Y - r.Size.Height/2,
			},
			PointTwo: commondata.Position{
				X: r.Position.X + r.Size.Width/2,
				Y: r.Position.Y + r.Size.Height/2,
			},
		},
		{
			PointOne: commondata.Position{
				X: r.Position.X + r.Size.Width/2,
				Y: r.Position.Y + r.Size.Height/2,
			},
			PointTwo: commondata.Position{
				X: r.Position.X - r.Size.Width/2,
				Y: r.Position.Y + r.Size.Height/2,
			},
		},
		{
			PointOne: commondata.Position{
				X: r.Position.X - r.Size.Width/2,
				Y: r.Position.Y + r.Size.Height/2,
			},
			PointTwo: commondata.Position{
				X: r.Position.X - r.Size.Width/2,
				Y: r.Position.Y - r.Size.Height/2,
			},
		},
	}
}
