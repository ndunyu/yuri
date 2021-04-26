package yuri

import (
	"log"
	"sync"
)

///Will queue jobs to be executed later

// Process send to be enqueued
type Process interface {
	Process(interface{})
}
type Queue struct {
	///Number of workers(consumers)goroutines to open
	Workers int
	///how many items will be held in the queue
	Capacity int
	///this is
	JobQueue chan interface{}
	Wg       *sync.WaitGroup
	///timeout in seconds
	///TimeOut  time.Duration
	//QuitChan will be used to close this worker
	//terminate the goroutines jobs
	QuitChan chan struct{}
	//JobCallBack is the function to be called when a job is received
	//it should implement Process
	//when a job is received what should happen ?(call back)
	JobCallBack Process
}

// NewWorker create a new job queue
func NewWorker(workers int, capacity int, jobCallBack Process) Queue {
	var wg sync.WaitGroup
	jobQueue := make(chan interface{}, capacity)
	quit := make(chan struct{})
	return Queue{
		Workers:     workers,
		JobQueue:    jobQueue,
		JobCallBack: jobCallBack,
		Wg:          &wg,
		QuitChan:    quit,
	}

}

// Stop close or the running goroutines
// and stops processing any more jobs
func (q *Queue) Stop() {


	close(q.QuitChan)
	//if &q.QuitChan!=nil {
	//	close(q.QuitChan)
	//}

}

// EnqueueJob use this to queue the jobs you need to execute
//Note you set the buffer size with capacity
//Returns false if the buffer is full
///else if it is accepted it returns true
func (q *Queue) EnqueueJob(job interface{}) bool {
	select {
	case q.JobQueue <- job:
		return true
	default:
		return false
	}
}

// StartWorkers start the workers add add them to wait group
func (q *Queue) StartWorkers() {
	for i := 0; i < q.Workers; i++ {
		q.Wg.Add(1)
		go q.worker()
	}
	q.Wg.Wait()
}

//each goroutine runs this
func (q *Queue) worker() {
	defer q.Wg.Done()
	for {
		select {
		case <-q.QuitChan:

			log.Println("closing the  workers")
			return

		case job := <-q.JobQueue:

			q.JobCallBack.Process(job)
		}
	}
}
