import { Event } from './event';

export interface Publisher {
  publish(...events: Event[]): Promise<void>;
}
