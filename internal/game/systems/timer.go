package systems

import rl "github.com/gen2brain/raylib-go/raylib"

type ITimerEntity interface {
	Init()
	Update()
	Draw()
	Done()
}

type TimerEntity struct {
	doneFunction func()
}

func (te *TimerEntity) Init()   {}
func (te *TimerEntity) Update() {}
func (te *TimerEntity) Draw()   {}
func (te *TimerEntity) Done() {
	te.doneFunction()
}

type Timer struct {
	timers []ITimerEntity
	times  []float64
}

func (s *Timer) Init() {
	s.timers = make([]ITimerEntity, 0)
}

func (s *Timer) Update() {
	// backwards because we might need to remove timers while iterating
	for i := len(s.timers) - 1; i >= 0; i-- {
		s.timers[i].Update()
		s.times[i] -= float64(rl.GetFrameTime())
		if s.times[i] <= 0 {
			s.timers[i].Done()
			s.RemoveTimerByIndex(i)
		}
	}
}

func (s *Timer) Draw() {
	for _, entity := range s.timers {
		entity.Draw()
	}
}

func (s *Timer) AddTimer(timer ITimerEntity, time float64) {
	s.timers = append(s.timers, timer)
	s.times = append(s.times, time)
}

func (s *Timer) RemoveTimer(entity ITimerEntity) {
	var index, found = 0, false
	for i, e := range s.timers {
		if e == entity {
			found = true
			index = i
			break
		}
	}
	if found {
		s.RemoveTimerByIndex(index)
	}
}

func (s *Timer) RemoveTimerByIndex(index int) {
	s.timers = append(s.timers[:index], s.timers[index+1:]...)
	s.times = append(s.times[:index], s.times[index+1:]...)
}

func (s *Timer) After(time float64, fn func()) {
	var t = new(TimerEntity)
	t.doneFunction = fn
	s.AddTimer(t, time)
}
