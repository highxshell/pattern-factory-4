package observer

import "fmt"

type Manager interface {
	Register(o Observer)
	DeRegister(o Observer)
	NotifyAll()
}

type ItemManager struct {
	item Item
}

func NewItemManager(item *Item) *ItemManager {
	return &ItemManager{
		item: *item,
	}
}

func (i *ItemManager) Register(o Observer) {
	i.item.observerList = append(i.item.observerList, o)
}

func (i *ItemManager) DeRegister(o Observer) {
	i.item.observerList = removeFromSlice(i.item.observerList, o)
}

func (i *ItemManager) NotifyAll() {
	for _, observer := range i.item.observerList {
		observer.update(i.item.name)
	}
}

func (i *ItemManager) UpdateAvailability() {
	fmt.Printf("Item %s is now in stock\n", i.item.name)
	i.item.inStock = true
	i.NotifyAll()
}

func removeFromSlice(observerList []Observer, observerToRemove Observer) []Observer {
	observerListLength := len(observerList)
	for i, observer := range observerList {
		if observerToRemove.getID() == observer.getID() {
			observerList[observerListLength-1], observerList[i] = observerList[i], observerList[observerListLength-1]
			return observerList[:observerListLength-1]
		}
	}
	return observerList
}
