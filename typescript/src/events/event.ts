import { v4 as uuid } from 'uuid';
import { Buffer } from 'buffer';

export const ErrInvalidEvent = new Error('invalid event');
export const ErrPayloadSerialization = new Error('event.payload_serialization');
export const ErrInternal = new Error('internal event');

// Event
export class Event {
  private id: string;
  private entityId: string;
  private topic: string;
  private payload: Uint8Array;
  private timestamp: Date;

  constructor(
    id: string,
    entityId: string,
    topic: string,
    payload: Uint8Array,
    timestamp: Date,
  ) {
    if (id.length === 0) {
      throw ErrInvalidEvent;
    }

    if (entityId.length === 0) {
      throw ErrInvalidEvent;
    }

    if (topic.length === 0) {
      throw ErrInvalidEvent;
    }

    if (!payload || payload.length === 0) {
      throw ErrInvalidEvent;
    }

    this.id = id;
    this.entityId = entityId;
    this.topic = topic;
    this.payload = payload;
    this.timestamp = timestamp;
  }

  static create(entityId: string, topic: string, payload: unknown): Event {
    let buff;
    try {
      const json = JSON.stringify(payload);
      buff = Buffer.from(json);
    } catch (err) {
      throw ErrPayloadSerialization;
    }

    return new Event(uuid(), entityId, topic, new Uint8Array(buff), new Date());
  }

  getId(): string {
    return this.id;
  }

  getEntityId(): string {
    return this.entityId;
  }

  getTopic(): string {
    return this.topic;
  }

  getPayload(): Uint8Array {
    return this.payload;
  }

  deserializePayload<T>(): T {
    const buff = Buffer.from(this.payload);

    try {
      return JSON.parse(buff.toString());
    } catch (err) {
      throw ErrPayloadSerialization;
    }
  }

  getTimestamp(): Date {
    return this.timestamp;
  }
}
