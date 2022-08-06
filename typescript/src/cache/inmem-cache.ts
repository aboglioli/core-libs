import { Cache } from './cache';

type Key = string | number | symbol;

interface Map<V> {
  [key: Key]: V;
}

export class InMemCache<V> implements Cache<Key, V> {
  private items: Map<V>;

  constructor() {
    this.items = {};
  }

  async all(): Promise<Map<V>> {
    return this.items;
  }

  async get(k: Key): Promise<V | null> {
    return this.items[k] || null;
  }

  async set(k: Key, v: V): Promise<void> {
    this.items[k] = v;
  }

  async delete(k: Key): Promise<void> {
    delete this.items[k];
  }
}
