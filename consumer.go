package Observer

type Consumer interface {
	Update(pubName string) 
	GetName() string 
}