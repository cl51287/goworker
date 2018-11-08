package goworker

import (
	"fmt"
)

// WorkShop is the main process
type WorkShop struct {
	workerManager *WorkerManager
	producers     []Producer
}

// SetWorkerManager is workerManager setter
func (workShop *WorkShop) SetWorkerManager(workManager *WorkerManager) {
	workShop.workerManager = workManager
}

// AddProducer is the producer add method
func (workShop *WorkShop) AddProducer(producer *Producer) {
}

// Run is ready to work
func (workShop *WorkShop) Run() {
	workShop.startWorkProcess()

	workShop.startProduceProcess()
}

func (workShop *WorkShop) startProduceProcess() {
	// start n goroutines and let them keep producing
	for i := 0; i < len(workShop.producers); i++ {
		// start a goroutine and let it keep producing
		go func(workerManager *WorkerManager, producer Producer) {
			for true {
				tasks := producer.Produce()
				for _, task := range tasks {
					err := workerManager.AddTask(task)
					if err != nil {
						fmt.Printf("Add Task error: %s\n", err)
					}
				}
			}
		}(workShop.workerManager, workShop.producers[i])
	}
}

func (workShop *WorkShop) startWorkProcess() {
	workShop.workerManager.StartWorking()
}
