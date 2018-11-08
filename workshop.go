package goworker

import (
	"time"
)

// WorkShop is the main process
type WorkShop struct {
	workerManager *WorkerManager
	producers     []Producer
}

func NewWorkShop(workerManager *WorkerManager, producers []Producer) *WorkShop {
	return &WorkShop{
		workerManager: workerManager,
		producers:     producers,
	}
}

// SetWorkerManager is workerManager setter
func (workShop *WorkShop) SetWorkerManager(workManager *WorkerManager) {
	workShop.workerManager = workManager
}

// AddProducer is the producer add method
func (workShop *WorkShop) AddProducer(producer Producer) {
	workShop.producers = append(workShop.producers, producer)
}

func (workShop *WorkShop) SetProducers(producers []Producer) {
	workShop.producers = producers
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

	workShop.workerManager.StartWorking()
}
