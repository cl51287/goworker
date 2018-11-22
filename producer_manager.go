package goworker

type ProducerManager struct {
	producers map[string]Producer
	agent     Agent
}

func (pm *ProducerManager) SetAgent(agent Agent) {
	pm.agent = agent
}

func (pm *ProducerManager) ReceiveTask(task *Task) {

}

func (pm *ProducerManager) AddProducer(producerName string, producer Producer) bool {
	pm.producers[producerName] = producer

	return true
}

func (pm *ProducerManager) GetProducer(producerName string) Producer {
	producer, ok := pm.producers[producerName]

	if ok {
		return producer
	}
	return nil
}

func (pm *ProducerManager) HasProducer(producerName string) bool {
	_, ok := pm.producers[producerName]

	return ok
}

func (pm *ProducerManager) DelProducer(producerName string) bool {

	return true
}

func (pm *ProducerManager) StopProducer(producerName string) bool {

	return true
}

func (pm *ProducerManager) StartProducer(producerName string) bool {

	return true
}

func (pm *ProducerManager) Start() bool {

	return true
}
