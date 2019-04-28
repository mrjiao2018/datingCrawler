package scheduler

import "crawler/engine"

// 简单 Request 调度器的实现
type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) ConfigureMasterWorkerChan(in chan engine.Request) {
	s.workerChan = in
}

func (s *SimpleScheduler) Submit(r engine.Request) {
	// 不开 goroutine 会在 ConcurrentEngine.createWorker()部分陷入死循环(循环等待)
	go func() { s.workerChan <- r }()
}
