package main

import (
	"crawler/engine"
	"crawler/persist"
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
	itemChan, err := persist.ItemSaver("dating_profile")
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueueScheduler{},
		WorkerCount: 100,
		ItemChan:    itemChan,
	}

	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})

}
