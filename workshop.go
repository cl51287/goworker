package goworker

import (
	"time"
)

// WorkShop is the main process
type WorkShop struct {
	workerManager   *WorkerManager
	producerManager *ProducerManager
}

func NewWorkShop() *WorkShop {
	return &WorkShop{
		workerManager: workerManager,
		producers:     producers,
	}
}

func (ws *WorkShop) AddProducer(producerName string, producer Producer) bool {

}

func (ws *WorkShop) AddStoppedProducer(producerName string, producer Producer) bool {

}

func (ws *WorkShop) DelProducer(producerName string) bool {

}

func (ws *WorkShop) StartProducer(producerName string) bool {

}

func (ws *WorkShop) StopProducer(producerName string) bool {

}

func (ws *WorkShop) AddWorker(workerName string, worker Worker) bool {

}

func (ws *WorkShop) AddStoppedWorker(workerName string, worker Worker) bool {

}

func (ws *WorkShop) DelWorker(workerName string) bool {

}

func (ws *WorkShop) StartWorker(workerName string) bool {

}

func (ws *WorkShop) StopWorker(workerName string) bool {

}

// SetWorkerManager is workerManager setter
func (workShop *WorkShop) SetWorkerManager(workManager *WorkerManager) {
	workShop.workerManager = workManager
}

// Run is ready to work
func (workShop *WorkShop) Run() {
	workShop.startWorkProcess()

	workShop.startProduceProcess()

	for true {
		time.Sleep(time.Second * 10)
	}
}

func (workShop *WorkShop) startProduceProcess() {
	if len(workShop.producers) == 0 {
		panic("producers not set")
	}

	// start n goroutines and let them keep producing
	for i := 0; i < len(workShop.producers); i++ {
		// start a goroutine and let it keep producing
		go func(workerManager *WorkerManager, producer Producer) {
			for true {
				tasks := producer.Produce()
				for _, task := range tasks {
					workerManager.AddTask(task)
				}
			}
		}(workShop.workerManager, workShop.producers[i])
	}
}

func (workShop *WorkShop) startWorkProcess() {
	if workShop.workerManager == nil {
		panic("worker manager not set")
	}

	go workShop.workerManager.StartWorking()
}
