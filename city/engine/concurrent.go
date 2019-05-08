package engine

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
	ItemChan    chan interface{}
}
type Scheduler interface {
	Submit(Request)
	//ConfigMushWorkerChan(chan Request) 我的queued没有实现这个方法,
	WorkerReady(chan Request)
	Run()
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	//in := make(chan Request)
	out := make(chan ParseResult)
	e.Scheduler.Run() //run进去先创建两个chan
	//e.Scheduler.ConfigMushWorkerChan(in) //将request放到workerchan里面  in就是workerchan 就是将request放到in

	for i := 0; i < e.WorkerCount; i++ {
		createWorker(out, e.Scheduler)
	}

	for _, r := range seeds {
		e.Scheduler.Submit(r) //将result提交
	}
	//out取数据 死循环 不断的输出,只要有数据 就输出

	for {
		result := <-out //从out里面源源不断的获取request
		for _, item := range result.Items {
			go func() { e.ItemChan <- item }()
		}
		//将你out 里面的request 发给scheduler
		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}

}

//并发了fetcher
func createWorker(out chan ParseResult, s Scheduler) {
	in := make(chan Request)
	go func() {
		for {
			//tell scheduler i am ready
			s.WorkerReady(in)
			request := <-in                //将in里面的url传递给fetcher
			result, err := worker(request) //fetch 里面的request 放到[]Request
			if err != nil {
				continue
			}
			out <- result // 再将返回的request存储到 out里面
		}
	}()
}
