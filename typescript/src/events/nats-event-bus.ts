import { Codec, JSONCodec, NatsConnection, Subscription as NatsSubscription } from 'nats';
import { Buffer } from 'buffer';

import { Event, Handler, Publisher, Subscriber } from '../events';

export interface NatsEvent {
  id: string;
  entity_id: string;
  topic: string;
  payload: Uint8Array | string;
  timestamp: Date;
}

export class NatsEventBus implements Publisher, Subscriber {
  private consumerGroup: string;
  private conn: NatsConnection;
  private codec: Codec<NatsEvent>;

  constructor(consumerGroup: string, conn: NatsConnection) {
    this.consumerGroup = consumerGroup;
    this.conn = conn;
    this.codec = JSONCodec();
  }

  async publish(...events: Event[]): Promise<void> {
    for (const event of events) {
      const natsEvent: NatsEvent = {
        id: event.getId(),
        entity_id: event.getEntityId(),
        topic: event.getTopic(),
        payload: event.getPayload(),
        timestamp: event.getTimestamp(),
      };

      const msgJson = JSON.stringify(natsEvent);
      const msgBuff = Buffer.from(msgJson);

      this.conn.publish(event.getTopic(), new Uint8Array(msgBuff));
    }
  }

  async subscribe(subject: string, handler: Handler): Promise<void> {
    const sub = this.conn.subscribe(subject, {
      queue: this.consumerGroup,
    });

    (async (sub: NatsSubscription) => {
      for await (const msg of sub) {
        const natsEvent = this.codec.decode(msg.data);

        const event = new Event(
          natsEvent.id,
          natsEvent.entity_id,
          natsEvent.topic,
          typeof natsEvent.payload === 'string'
            ? new Uint8Array(Buffer.from(natsEvent.payload as string, 'base64'))
            : natsEvent.payload,
          natsEvent.timestamp,
        );

        await handler.handle(event);
      }
    })(sub);
  }
}
