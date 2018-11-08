package main

import (
	"fmt"
	"goworker"
	"math/rand"
)

type Producer struct {
}

func (producer *Producer) Produce() []*Task {
	tasks := make([]*Task, 10)
	for i := 0; i < len(tasks); i++ {
		tasks[i].ID = i
	}
	return tasks
}

type Worker struct {
	ID int
}

func (worker *Worker) Handle(task *Task) {
	fmt.Println("worker id is %d, task id is %d.\n", worker.ID, task.ID)
}

type Dispatcher struct {
}

func (dispatcher *Dispatcher) dispatch(workers []*Worker) *Worker {
	i := rand.Intn(len(Worker)) - 1
	return workers[i]
}

func main() {
	producers := make([]*Producer, 2)
	for i := 0; i < len(producers); i++ {
		producers[i] = &Producer{}
	}

	dispatcher := &Dispatcher{}

	workers := make([]*Worker, 5)
	for j := 0; j < len(producers); j++ {
		workers[j] = &Worker{ID: j}
	}
	workerManager := goworker.NewWorkerManager(dispatcher, workers)
	workShop := goworker.NewWorkShop(workerManager, producers)
	workShop.Run()
}
