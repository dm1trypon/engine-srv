package effect

/*
Effect - contains data about positive and negative effects.
	- Data - contains:
		- kind <string> - kind of effect.
		- duration <int> - duration.
		- chance <int> - chance of triggering.
		- speed <int> - possible additional speed, default is 0.
		- rateOfFire <int> - possible rate of fire, default is 0.
*/
type Effect struct {
	kind       string
	chance     int
	duration   int
	speed      int
	rateOfFire int
}
