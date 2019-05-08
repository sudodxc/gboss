package scheduler

import (
	"city/engine"
)

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) ConfigMushWorkerChan(c chan engine.Request) {
	s.workerChan = c
}
func (s *SimpleScheduler) Submit(r engine.Request) {
	go func() { s.workerChan <- r }() //因为前面放的速度非常的快,如果这里不开协程 会死掉
}
