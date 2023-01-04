package main

import "fmt"

type Topic interface {
	register(observer Observer)
	broadcast()
}

type Observer interface {
	GetId() string
	UpdateValue(string)
}

type Item struct {
	observers []Observer
	name      string
	available bool
}

func NewItem(name string) *Item {
	return &Item{
		name: name,
	}
}

func (i *Item) Register(observer Observer) {
	i.observers = append(i.observers, observer)
}

func (i *Item) Broadcast() {
	for _, observer := range i.observers {
		observer.UpdateValue(i.name)
	}
}

func (i *Item) UpdateAvailable() {
	fmt.Printf("Item %s is available\n", i.name)
	i.available = true
	i.Broadcast()
}

type EmailCLient struct {
	id string
}

func (ec *EmailCLient) GetId() string {
	return ec.id
}

func (ec *EmailCLient) UpdateValue(value string) {
	fmt.Printf("Sending Email - %s available from client %s\n", value, ec.id)
}

func main() {
	nvidiaItem := NewItem("RTX 3080")
	firstObserver := &EmailCLient{
		id: "abc12345",
	}
	secondObserver := &EmailCLient{
		id: "xyz6789",
	}

	nvidiaItem.Register(firstObserver)
	nvidiaItem.Register(secondObserver)
	nvidiaItem.UpdateAvailable()

}
