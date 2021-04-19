package yuri

import (
	"log"
	"sync"
)

///Will queue jobs to be executed later

///send to be enqueued
type Processv1 func(interface{})
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
	QuitChan chan struct{}
	///when a job is received what should happen ?(call back)
	JobCallBack  Process
}

func NewWorker(workers int,capacity int,jobCallBack Process) Queue {
	var wg sync.WaitGroup
	jobQueue := make(chan interface{},capacity )
	//cancelChan := make(chan bool)
	quit := make(chan struct{})
	return Queue{
		Workers:  workers,
		JobQueue: jobQueue,
		JobCallBack: jobCallBack,
		Wg:       &wg,
		QuitChan: quit,
	}

}

//close or the running goroutines
func (q *Queue) Stop() {
	log.Println("called here")

	close(q.QuitChan)
	//if &q.QuitChan!=nil {
	//	close(q.QuitChan)
	//}


}


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

///start the workers add add them to wait group
func (q *Queue) StartWorkers() {
	for i := 0; i < q.Workers; i++ {

		q.Wg.Add(1)
		go q.worker()
		///go q.worker(jobChan)
	}
	log.Println("am waiting")
	///q.Wg.Wait()
	log.Println("finished")
}

//each goroutine runs this
func (q *Queue) worker() {
	defer q.Wg.Done()
	for {
		select {
		case  <-q.QuitChan:

			log.Println("its closing me")
			return

		case job := <-q.JobQueue:

			q.JobCallBack.Process(job)
		}
	}
}
///TryEnqueue(job Job, jobChan <-chan Job)

///TryEnqueue(job Job, jobChan chan<- Job)
