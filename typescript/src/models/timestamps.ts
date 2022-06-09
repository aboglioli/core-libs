import { Error } from '../errors';
import { Metadata } from '../collections';

export const ErrInvalidTimestamps = Error.define('timestamps.invalid');

export class Timestamps {
  private createdAt: Date;
  private updatedAt: Date;
  private deletedAt?: Date;

  constructor(createdAt: Date, updatedAt: Date, deletedAt?: Date) {
    const m = Metadata.with('created_at', createdAt)
      .and('updated_at', updatedAt)
      .and('deleted_at', deletedAt);

    if (updatedAt < createdAt) {
      throw Error.create(ErrInvalidTimestamps, 'update date is before create date', m);
    }

    if (deletedAt) {
      if (deletedAt < createdAt) {
        throw Error.create(ErrInvalidTimestamps, 'delete date is before create date', m);
      }

      if (deletedAt < updatedAt) {
        throw Error.create(ErrInvalidTimestamps, 'delete date is before update date', m);
      }
    }

    this.createdAt = createdAt;
    this.updatedAt = updatedAt;
    this.deletedAt = deletedAt;
  }

  static create(): Timestamps {
    const now = new Date();
    return new Timestamps(now, now);
  }

  getCreatedAt(): Date {
    return this.createdAt;
  }

  getUpdatedAt(): Date {
    return this.updatedAt;
  }

  update(): Timestamps {
    return new Timestamps(this.createdAt, new Date(), this.deletedAt);
  }

  getDeletedAt(): Date | undefined {
    return this.deletedAt;
  }

  delete(): Timestamps {
    return new Timestamps(this.createdAt, this.updatedAt, new Date());
  }
}
