import { v4 as uuid } from 'uuid';
import { Buffer } from 'buffer';

import { Error } from '../errors';
import { Metadata } from '../collections';

export const ErrInvalidEvent = Error.define('event.invalid');
export const ErrEventPayloadSerialization = Error.define('event.payload_serialization');
export const ErrEventInternal = Error.define('event.internal');

// Publisher and subscriber
export interface Publisher {
  publish(...events: Event[]): Promise<void>;
}

export interface Handler {
  handle(event: Event): Promise<void>;
}

export interface Subscriber {
  subscribe(subject: string, handler: Handler): Promise<void>;
}

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
    const m = Metadata.with('id', id)
      .and('entity_id', entityId)
      .and('topic', topic)
      .and('payload', payload)
      .and('timestamp', timestamp);

    if (id.length === 0) {
      throw Error.create(ErrInvalidEvent, 'payload id is empty', m);
    }

    if (entityId.length === 0) {
      throw Error.create(ErrInvalidEvent, 'payload entity_id is empty', m);
    }

    if (topic.length === 0) {
      throw Error.create(ErrInvalidEvent, 'payload topic is empty', m);
    }

    if (!payload || payload.length === 0) {
      throw Error.create(ErrInvalidEvent, 'payload is empty', m);
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
      throw Error.wrap(
        ErrEventPayloadSerialization,
        err,
        'could not serialize event payload',
      );
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
      throw Error.wrap(
        ErrEventPayloadSerialization,
        err,
        'could not deserialize event payload',
      );
    }
  }

  getTimestamp(): Date {
    return this.timestamp;
  }
}
