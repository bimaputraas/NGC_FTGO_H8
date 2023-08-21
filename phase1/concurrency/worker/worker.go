package worker

func (wi *WorkerInstance) Do(job func()) {
	wi.jobsWaitGroup.Add(1)
	wi.jobsChannel <- job
}

func (wi *WorkerInstance) Wait() {
	wi.jobsWaitGroup.Wait()
}