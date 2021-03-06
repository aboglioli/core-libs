import { Event } from './event';

export interface Publishable {
  entityId(): string;
  topic(): string;
}

export class EventCollector {
  private events: Event[];

  constructor(events: Event[]) {
    this.events = events;
  }

  static create() {
    return new EventCollector([]);
  }

  record(p: Publishable) {
    const event = Event.create(p.entityId(), p.topic(), p);

    this.events.push(event);
  }

  all(): Event[] {
    return this.events;
  }

  drain(): Event[] {
    const events = this.events;

    this.events = [];

    return events;
  }
}
