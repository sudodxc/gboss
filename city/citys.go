package main

import (

	"city/boss/parse"
	"city/engine"
	"city/persist"
	"city/scheduler"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 10, //太大会超过网站的最大访问限制 可以通过time.Tick/等手段限制一下,我图省事就减少了workerCount
		ItemChan:    persist.ItemSaver(),
	}

	e.Run(engine.Request{
		Url:       "https://www.zhipin.com/",
		ParseFunc: parse.ParseCityList,
	})

}
