package commondata

/*
Size - contains the width and height of the object.
	- Data - contains:
		- Width <int> - width of the object.
		- Height <int> - height of the object.
*/
type Size struct {
	Width  int
	Height int
}

/*
Speed - contains the speed data for the x and y axis, as well as the maximum speed.
	- Data - contains:
		- X <int> - X-axis speed.
		- Y <int> - Y-axis speed.
		- Max <int> - limit of speed.
*/
type Speed struct {
	X   int
	Y   int
	Max int
}

/*
Position - data that contains the coordinates of the object.
	- Data - contains:
		- X <int> - X-axis position.
		- Y <int> - Y-axis position.
*/
type Position struct {
	X int
	Y int
}
