package engine

import (
	"log"
)

type SingleThreadEngine struct {
}

func (SingleThreadEngine) Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	//此处是将 requests 当做一个任务队列，从网页里面不断爬取新的url，并将这些 url 入队
	for len(requests) > 0 {
		//1. 先从 requests 队列中取出一个任务
		r := requests[0]
		requests = requests[1:]

		//2. 开启 worker
		parseResult, err := Worker(r)
		if err != nil {
			continue
		}

		//3. 将 parserResult 中新增的 url 放入 requests 队列中
		requests = append(requests, parseResult.Requests...)

		for _, item := range parseResult.Items {
			log.Printf("Got %v", item)
		}
	}
}
