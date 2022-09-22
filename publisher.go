package Observer

type Publisher struct {
	Name string 
	Consumers map[string]Consumer
}

func (pub *Publisher) Subscribe(consumer Consumer) {
	pub.Consumer[consumer.GetName()] = consumer
}

func (pub *Publisher) UnSubscribe(consumer Consumer) {
	delete(pub.Consumer, consumer.GetName())
}


func (pub *Publisher) Notify() {
	for _, consumer := range pub.Consumers {
		consumer.Update(pub.Name) 
	}
}