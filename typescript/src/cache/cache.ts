import { Error } from '../errors';

export const ErrCacheInternal = Error.define('cache.internal');

export interface Cache<K, V> {
  get(k: K): Promise<V | null>;
  set(k: K, v: V): Promise<void>;
  delete(k: K): Promise<void>;
}
