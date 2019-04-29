package main

import (
	"crawler/engine"
	"crawler/scheduler"
	"crawler/zhenai/parser"
)

// 单线程版本爬虫
//func main() {
//	engine.SingleThreadEngine{}.Run(engine.Request{
//		Url:        "http://www.zhenai.com/zhenghun",
//		ParserFunc: parser.ParseCityList,
//	})
//}

// 多线程版本爬虫
func main() {
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueueScheduler{},
		WorkerCount: 100,
	}

	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})

}
