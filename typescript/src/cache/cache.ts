export const ErrCacheInternal = new Error('internal cache error');

export interface Cache<K, V> {
  get(k: K): Promise<V | null>;
  set(k: K, v: V): Promise<void>;
  delete(k: K): Promise<void>;
}
