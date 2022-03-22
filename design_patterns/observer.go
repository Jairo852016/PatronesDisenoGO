package main

import "fmt"

type Topic interface {
	register(observer Observer)
	brosdcast()
}

type Observer interface {
	getId() string
	updateValue(string)
}

// Item -> No Disponible
// Item -> Avise -> HAY ITEM PS5

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
func (i *Item) UpdateAvailable() {
	fmt.Printf("Item %d is available \n", i.name)
	i.available = true
	i.brosdcast()
}
func (i *Item) brosdcast() {
	for _, observer := range i.observers {
		observer.updateValue(i.name)
	}
}
func (i *Item) register(Observer Observer) {
	i.observers = append(i.observers, Observer)
}

type EmailClient struct {
	id string
}

func (eC *EmailClient) updateValue(value string) {
	fmt.Printf("Sending Email - %s available from cliente %s\n", value, eC.id)
}
func (eC *EmailClient) getId() string {
	return eC.id

}
func main() {
	nvidiaItem := NewItem("Rtx 3080")
	fisrtObserver := &EmailClient{
		id: "12ab",
	}
	secondObserver := &EmailClient{
		id: "34dc",
	}
	nvidiaItem.register(fisrtObserver)
	nvidiaItem.register(secondObserver)
	nvidiaItem.UpdateAvailable()
}
