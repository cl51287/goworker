package main

import (
	"fmt"
	"math/rand"

	"../../goworker"
)

type Producer struct {
}

func (producer *Producer) Produce() []*goworker.Task {
	tasks := make([]*goworker.Task, 10)
	for i := 0; i < len(tasks); i++ {
		tasks[i] = &goworker.Task{
			ID: int32(i),
		}
	}
	return tasks
}

type Worker struct {
	ID int
}

func (worker *Worker) Handle(task *goworker.Task) error {
	fmt.Printf("worker id is %d, task id is %d.\n", worker.ID, task.ID)
	return nil
}

type Dispatcher struct {
}

func (dispatcher *Dispatcher) Dispatch(workers []goworker.Worker) goworker.Worker {
	i := rand.Intn(len(workers) - 1)
	return workers[i]
}

func main() {
	producers := make([]goworker.Producer, 2)
	for i := 0; i < len(producers); i++ {
		producers[i] = &Producer{}
	}

	dispatcher := &Dispatcher{}

	workers := make([]goworker.Worker, 5)
	for j := 0; j < len(workers); j++ {
		workers[j] = &Worker{ID: j}
	}
	workerManager := goworker.NewWorkerManager(dispatcher, workers)
	workShop := goworker.NewWorkShop(workerManager, producers)
	workShop.Run()
}
