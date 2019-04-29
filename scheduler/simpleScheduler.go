package scheduler

import "crawler/engine"

// 简单 Request 调度器的实现，所有 worker 公用一个 channel
type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) WorkerChan() chan engine.Request {
	return s.workerChan
}

func (s *SimpleScheduler) WorkerReady(chan engine.Request) {

}

func (s *SimpleScheduler) Run() {
	s.workerChan = make(chan engine.Request)
}

func (s *SimpleScheduler) Submit(r engine.Request) {
	// 不开 goroutine 会在 ConcurrentEngine.createWorker()部分陷入死循环(循环等待)
	// 该种方式无法控制将 request 分发给哪一个具体的 worker
	go func() { s.workerChan <- r }()
}
