package effect

func (e *Effect) Create() {
	e = &Effect{
		kind:           "",
		speed:          0,
		isInvisibility: false,
		duration:       0,
	}
}

func (e *Effect) GetKind() string {
	return e.kind
}

func (e *Effect) GetSpeed() int {
	return e.speed
}

func (e *Effect) GetIsInvisibility() bool {
	return e.isInvisibility
}

func (e *Effect) GetDuration() int {
	return e.duration
}

func (e *Effect) SetKind(kind string) {
	e.kind = kind
}

func (e *Effect) SetSpeed(speed int) {
	e.speed = speed
}

func (e *Effect) SetIsInvisibility(isInvisibility bool) {
	e.isInvisibility = isInvisibility
}

func (e *Effect) SetDuration(duration int) {
	e.duration = duration
}
