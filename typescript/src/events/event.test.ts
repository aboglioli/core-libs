import { Event } from './event';

describe('Event', () => {
  test('create event', () => {
    const event = Event.create('entity#01', 'topic.code', {
      msg: 'Hello World',
    });
    expect(event.getEntityId()).toBe('entity#01');
    expect(event.getTopic()).toBe('topic.code');
    expect(event.getPayload()).toBeInstanceOf(Uint8Array);
  });

  test('payload serialization and deserialization', () => {
    const event = Event.create('entity#01', 'topic.code', {
      msg: 'Hello World',
    });
    const payload = event.deserializePayload();
    expect(payload).toEqual({ msg: 'Hello World' });
  });
});
