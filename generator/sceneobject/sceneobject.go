package sceneobject

func (s *SceneObject) Create() *SceneObject {
	s = &SceneObject{
		maxCount:  0,
		respDelay: 0,
		kind:      "",
	}

	return s
}

func (s *SceneObject) GetMaxCount() int {
	return s.maxCount
}

func (s *SceneObject) GetRespDelay() int {
	return s.respDelay
}

func (s *SceneObject) GetKind() string {
	return s.kind
}

func (s *SceneObject) SetMaxCount(maxCount int) {
	s.maxCount = maxCount
}

func (s *SceneObject) SetRespDelay(respDelay int) {
	s.respDelay = respDelay
}

func (s *SceneObject) SetKind(kind string) {
	s.kind = kind
}
