package collision

import "github.com/dm1trypon/engine-srv/goapi/models/commondata"

type CircCircPair struct {
	CircleOne Circle
	CircleTwo Circle
}

type CircRectPair struct {
	Circle    Circle
	Rectangle Rectangle
}

type Segment struct {
	PointOne commondata.Position
	PointTwo commondata.Position
}

type Rectangle struct {
	Size     commondata.Size
	Position commondata.Position
	Segments []Segment
}

type Circle struct {
	Radius   float64
	Position commondata.Position
}
