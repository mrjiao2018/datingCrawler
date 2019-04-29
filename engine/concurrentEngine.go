package engine

import (
	"log"
)

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

type Scheduler interface {
	Submit(Request)
	ConfigureMasterWorkerChan(chan Request)
	//WorkerReady(chan Request)
	//Run()
}

func (e *ConcurrentEngine) Run(seeds ...Request) {

	// 1. 设定所有 worker 共享的输入输出 channel
	// 并将 input channel 交给 Scheduler 调度器来管理
	in := make(chan Request)
	out := make(chan ParseResult)
	e.Scheduler.ConfigureMasterWorkerChan(in)

	// 3. 开创 workerCount 个 worker 来完成 scheduler 分发的任务
	for i := 0; i < e.WorkerCount; i++ {
		// 当前所有 worker 共用相同的 input channel 和 output channel
		createWorker(in, out)
	}

	// 2. 将所有 Request 交给 Scheduler 来进行调度分发
	for _, request := range seeds {
		e.Scheduler.Submit(request)
	}

	// 4. 接收 out 中的结果
	itemCount := 0
	for {
		result := <-out
		for _, item := range result.Items {
			log.Printf("Got item %d : %v\n", itemCount, item)
			itemCount++
		}

		// 5. 获得 Item 之后，要将 result 中的所有 request 再送回给 Scheduler
		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}
}

func createWorker(in chan Request, out chan ParseResult) {
	go func() {
		for {
			request := <-in
			result, err := Worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
