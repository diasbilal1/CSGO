package observer

type EventObserver interface {
	Update(event string)
}

type EventNotifier struct {
	Observers []EventObserver
}

func (n *EventNotifier) RegisterObserver(observer EventObserver) {
	n.Observers = append(n.Observers, observer)
}

func (n *EventNotifier) Notify(event string) {
	for _, observer := range n.Observers {
		observer.Update(event)
	}
}
