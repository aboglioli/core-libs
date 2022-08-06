package events

type Publishable interface {
	EntityId() string
	Topic() string
}

type Collector struct {
	events []*Event
}

func NewCollector() *Collector {
	return &Collector{
		events: make([]*Event, 0),
	}
}

func (c *Collector) Record(publishable Publishable) error {
	event, err := CreateEvent(publishable.EntityId(), publishable.Topic(), publishable)
	if err != nil {
		return err
	}

	c.events = append(c.events, event)

	return nil
}

func (c *Collector) Events() []*Event {
	return c.events
}

func (c *Collector) Drain() []*Event {
	events := c.events

	c.events = make([]*Event, 0)

	return events
}
