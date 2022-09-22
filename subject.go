package Observer

type Subject interface {
	Subscribe(Consumer)
	UnSubscribe(Consumer)
	Notify() 
}