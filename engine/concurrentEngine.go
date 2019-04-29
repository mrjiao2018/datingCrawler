package engine

import (
	"log"
)

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

type Scheduler interface {
	ReadyNotifier
	// Scheduler 获取 Request 任务
	Submit(Request)

	// Scheduler 给 worker 分配 input channel 的方式
	// SimpleScheduler 实现中，所有的 worker 共用相同的 input channel
	// QueueScheduler 实现中，每个 worker 会获得一个新的 input channel
	WorkerChan() chan Request

	// 任务调度机制，即如何将 request 分发到 worker 上
	// SimpleScheduler 实现中，所有的 worker 来抢 request，不加以调控，负载均衡难以控制
	// QueueScheduler 实现中，通过 request 队列和 worker 队列加以控制
	Run()
}

type ReadyNotifier interface {
	// 在外界确认有 worker 可用的情况下，向 worker 发送 request
	WorkerReady(chan Request)
}

// 记录已经爬取过的 user 的 url 和 city 的 url，防止重复爬取，
// 但同时要保证查询效率为 O(1)，否则后期百万数据量时，查询性能很可能成为瓶颈，
// 解决方案：hashMap，key 为 url，value 自定义，不为空即可
var visitedUrls = make(map[string]bool)

func (e *ConcurrentEngine) Run(seeds ...Request) {

	// 1. 设定所有 worker 共享的输出 channel
	out := make(chan ParseResult)
	e.Scheduler.Run()

	// 2. 开创 workerCount 个 worker 来完成 scheduler 分发的任务
	for i := 0; i < e.WorkerCount; i++ {
		// 当前所有 worker 共用相同的 output channel
		// 而 input channel 则由 Scheduler 来分配，不同的 Scheduler 的分配方式不一样：
		createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	// 3. 将所有 Request 交给 Scheduler 来进行调度分发
	for _, request := range seeds {
		if isDuplicated(request.Url) {
			continue
		}
		e.Scheduler.Submit(request)
	}

	// 4. 接收 out 中的结果
	itemCount := 0
	for {
		result := <-out
		for _, item := range result.Items {
			log.Printf("Got item %d : %+v\n", itemCount, item)
			itemCount++
		}

		// 5. 获得 Item 之后，要将 result 中的所有 request 再送回给 Scheduler
		for _, request := range result.Requests {
			if isDuplicated(request.Url) {
				continue
			}

			e.Scheduler.Submit(request)
		}
	}
}

func createWorker(in chan Request, out chan ParseResult, ready ReadyNotifier) {
	go func() {
		for {
			ready.WorkerReady(in)
			request := <-in
			result, err := Worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}

func isDuplicated(url string) bool {
	if visitedUrls[url] {
		return true
	}
	visitedUrls[url] = true
	return false
}
