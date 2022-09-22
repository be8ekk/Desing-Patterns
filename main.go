package main

import "simple/pkg" 

func main() {
	sub1 := &Observer.Subscriber{Name: "Sub-1"}; 
	sub2 := &Observer.Subscriber{Name: "Sub-2"}; 
	sub3 := &Observer.Subscriber{Name: "Sub-3"}; 
	subN := &Observer.Subscriber{Name: "Sub-N"}; 
	chanel := Observer.Publisher {
		Name: "Publisher channel", 
		Consumers: map[string]Observer.Consumer
	}
	chanel.Subscribe(sub1)
	chanel.Subscribe(sub2)
	chanel.Subscribe(sub3)
	chanel.Subscribe(subN)
	chanel.Notify() 
	println("Unsubscribe Sub-3") 
	chanel.Unsubscribe(sub3)
	chanel.Notify()f
}