package scheduler

import "crawler/engine"

// 队列 Request 调度器的实现，每个 worker 有一个自己的 channel
type QueueScheduler struct {
	requestChan chan engine.Request
	// workerChan 是一个 chan worker 类型，
	// 每一个 worker 对外提供的是 chan engine.Request
	workerChan chan (chan engine.Request)
}

func (s *QueueScheduler) WorkerChan() chan engine.Request {
	return make(chan engine.Request)
}

func (s *QueueScheduler) Submit(request engine.Request) {
	s.requestChan <- request
}

func (s *QueueScheduler) WorkerReady(worker chan engine.Request) {
	s.workerChan <- worker
}

// 队列调度的实现
// 维护 request 队列和 worker 队列，
// 当队列为空时，直接向队列中添加元素
// 队列不为空时，则取出队列中的队首元素，将 request 分发给 worker
func (s *QueueScheduler) Run() {
	s.requestChan = make(chan engine.Request)
	s.workerChan = make(chan chan engine.Request)

	go func() {
		var requestQ []engine.Request
		var workerQ []chan engine.Request

		for {
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeRequest = requestQ[0]
				activeWorker = workerQ[0]
			}

			select {
			case r := <-s.requestChan:
				requestQ = append(requestQ, r)
			case w := <-s.workerChan:
				workerQ = append(workerQ, w)
			case activeWorker <- activeRequest:
				requestQ = requestQ[1:]
				workerQ = workerQ[1:]
			}
		}
	}()
}
