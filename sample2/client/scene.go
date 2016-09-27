package main

type Scene struct {
	objects []Object
}

func NewScene() *Scene {
	return &Scene{
		objects: nil,
	}
}

func (s *Scene) Update() {
	for _, o := range s.objects {
		o.Update()
	}
}

//func (s *Scene) Start() {
//	go s.listen()
//}

//func (s *Scene) listen() {
//	ticker := time.NewTicker(time.Second / 60.0)

//	for {
//		select {

//		case <-ticker.C:
//			s.Update()

//		}
//	}
//}
