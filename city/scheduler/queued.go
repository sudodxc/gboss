package scheduler

import (
	"city/engine"
)

type QueuedScheduler struct {
	requestChan chan engine.Request      //request 的chan
	workerChan  chan chan engine.Request //将你前面所有的worker都灌在这个chan里面
}

func (s *QueuedScheduler) Submit(r engine.Request) { //当有人submit的时候就将request加进去
	s.requestChan <- r
}
func (s *QueuedScheduler) WorkerReady(w chan engine.Request) {
	s.workerChan <- w
}

func (s *QueuedScheduler)Run(){
	s.workerChan=make(chan chan engine.Request)
	s.requestChan=make(chan engine.Request) 
	go func() {
		//建立两个队列 :=后面要跟{}来说明里面是什么东西 所以这里用var
		var requestQ []engine.Request
		var workerQ []chan engine.Request
		for {
			//如果requestQ和workerQ里面都不为空
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeRequest = requestQ[0]
				activeWorker = workerQ[0]
			}
			select {
			//get request
			case r := <-s.requestChan:
				requestQ = append(requestQ, r) //我收到了request就让request排队
				//get workers 我想seed r to ?worker 不能直接这样
			case w := <-s.workerChan:
				//我想seed ?next_request to worker 不能直接这样
				workerQ = append(workerQ, w) //我收到了worker就让request排队
			case activeWorker <- activeRequest:
				requestQ = requestQ[1:]
				workerQ = workerQ[1:]
			}
		}

	}()

}

