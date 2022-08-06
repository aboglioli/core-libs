import { Event } from './event';

export interface Handler {
  handle(event: Event): Promise<void>;
}

export interface Subscriber {
  subscribe(subject: string, handler: Handler): Promise<void>;
}
