package effect

/*
Effect - effect data for Special runes.
Responsible for temporarily changing the characteristics of these objects.
	- Data - contains:
		- kind <string> - kind of effect.
		- speed <int> - variable speed of the object. Default value is 0.
		- isInvisibility <int> - possible invisibility, default is false.
		- duration <int> - duration of the effect.
*/
type Effect struct {
	kind           string
	speed          int
	isInvisibility bool
	duration       int
}
