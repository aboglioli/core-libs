import slug from 'slug';
import { v4 as uuid } from 'uuid';

import { Error } from '../errors';
import { Metadata } from '../collections';

export const ErrInvalidId = Error.define('id.invalid');

export class StrId {
  private id: string;

  constructor(id: string) {
    if (id.length == 0) {
      throw Error.create(ErrInvalidId, 'empty string id', Metadata.with('id', id));
    }

    this.id = id;
  }

  static generateUuid(): StrId {
    return new StrId(uuid());
  }

  static generateSlug(str: string): StrId {
    return new StrId(slug(str));
  }

  value(): string {
    return this.id;
  }

  equals(other: StrId): boolean {
    return this.id === other.id;
  }
}
