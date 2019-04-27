package engine

import (
	"crawler/fetcher"
	"log"
)

func Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	//此处是将 requests 当做一个任务队列，从网页里面不断爬取新的url，并将这些 url 入队
	for len(requests) > 0 {
		//1. 先从 requests 队列中取出一个任务
		r := requests[0]
		requests = requests[1:]

		//2. fetcher 将网页爬取下来，转成 utf-8 格式，获取 bytes []byte
		bytes, err := fetcher.Fetch(r.Url)
		log.Printf("Fetching url : %s", r.Url)
		if err != nil {
			log.Printf("Fetcher: error fetching url : %s, %v", r.Url, err)
			continue
		}

		//3. 由 ParserFunc 来解析 bytes，并将结果封装到 parserResult 中
		parserResult := r.ParserFunc(bytes)

		//4. 将 parserResult 中新增的 url 放入 requests 队列中
		requests = append(requests, parserResult.Requests...)

		for _, item := range parserResult.Items {
			log.Printf("Got %v", item)
		}
	}
}
