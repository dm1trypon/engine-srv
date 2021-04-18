package effect

func (e *Effect) Create() {
	e = &Effect{
		kind:       "",
		chance:     0,
		duration:   0,
		speed:      0,
		rateOfFire: 0,
	}
}

func (e *Effect) GetKind() string {
	return e.kind
}

func (e *Effect) GetChance() int {
	return e.chance
}

func (e *Effect) GetDuration() int {
	return e.duration
}

func (e *Effect) GetSpeed() int {
	return e.speed
}

func (e *Effect) GetRateOfFire() int {
	return e.rateOfFire
}

func (e *Effect) SetKind(kind string) {
	e.kind = kind
}

func (e *Effect) SetChance(chance int) {
	e.chance = chance
}

func (e *Effect) SetDuration(duration int) {
	e.duration = duration
}

func (e *Effect) SetSpeed(speed int) {
	e.speed = speed
}

func (e *Effect) SetRateOfFire(rateOfFire int) {
	e.rateOfFire = rateOfFire
}
