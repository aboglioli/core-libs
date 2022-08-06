import { Event } from './event';
import { Publisher } from './publisher';
import { Handler, Subscriber } from './subscriber';

interface Subscription {
  subject: string;
  handler: Handler;
}

function subjectHasTopic(subject: string, topic: string): boolean {
  if (subject === topic) {
    return true;
  }

  const subjectParts = subject.split('.');
  const topicParts = topic.split('.');

  if (subjectParts.length != topicParts.length) {
    return false;
  }

  return subjectParts.every((subjectPart, i) => {
    const topicPart = topicParts[i];

    return subjectPart === '*' || subjectPart === topicPart;
  });
}

export class LocalEventBus implements Publisher, Subscriber {
  private subscriptions: Subscription[];

  constructor() {
    this.subscriptions = [];
  }

  async publish(...events: Event[]): Promise<void> {
    for (const event of events) {
      for (const subscription of this.subscriptions) {
        if (subjectHasTopic(subscription.subject, event.getTopic())) {
          subscription.handler.handle(event);
        }
      }
    }
  }

  async subscribe(subject: string, handler: Handler): Promise<void> {
    this.subscriptions.push({ subject, handler });
  }
}
