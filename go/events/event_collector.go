package events

type Publishable interface {
	EntityId() string
	Topic() string
}

type EventCollector struct {
	events []*Event
}

func NewEventCollector(events []*Event) *EventCollector {
	return &EventCollector{
		events: events,
	}
}

func CreateEventCollector() *EventCollector {
	return NewEventCollector(make([]*Event, 0))
}

func (c *EventCollector) Record(p Publishable) error {
	event, err := CreateEvent(p.EntityId(), p.Topic(), p)
	if err != nil {
		return err
	}

	c.events = append(c.events, event)

	return nil
}

func (c *EventCollector) All() []*Event {
	return c.events
}

func (c *EventCollector) Drain() []*Event {
	events := c.events

	c.events = make([]*Event, 0)

	return events
}
