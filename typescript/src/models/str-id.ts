import slug from 'slug';
import { v4 as uuid } from 'uuid';

export const ErrInvalidId = new Error('invalid id');

export class StrId {
  private id: string;

  constructor(id: string) {
    if (id.length == 0) {
      throw ErrInvalidId;
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
