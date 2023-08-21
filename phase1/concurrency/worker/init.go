package worker

import (
	"errors"
	"sync"
)

type WorkerInstance struct {
	jobsChannel chan func()
	jobsWaitGroup *sync.WaitGroup
}

func NewWorkerInstance(numWorkers int) (*WorkerInstance, error) {
	if numWorkers == 0 {
		return nil, errors.New("zero worker")
	}
	wi := &WorkerInstance{
		jobsChannel: make(chan func()),
		jobsWaitGroup: &sync.WaitGroup{},
	}

	for i := 0; i < numWorkers; i++ {
		go func() {
			for job := range wi.jobsChannel {
				job()
				wi.jobsWaitGroup.Done()
			}
		}()
	}

	return wi, nil
}