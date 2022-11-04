package gworker

type workerManage struct {
	workers []Worker
	c       chan Worker
}

func Run(ws []Worker) {
	wm := &workerManage{
		workers: ws,
		c:       make(chan Worker),
	}
	wm.run()
}
func (w *workerManage) run() {
	for _, worker := range w.workers {
		go w.runWorker(worker)
	}
	for worker := range w.c {
		go w.runWorker(worker)
	}
}
func (w *workerManage) runWorker(worker Worker) {
	defer func() {
		v := recover()
		if v != nil {
			w.c <- worker
		}
	}()
	e := worker.Init()
	if e != nil {
		panic(e.Error())
	}
	e = worker.Execute()
	if e != nil {
		panic(e.Error())
	}
}
