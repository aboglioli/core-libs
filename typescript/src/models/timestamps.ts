export const ErrInvalidTimestamps = new Error('invalid timestamps');

export class Timestamps {
  private createdAt: Date;
  private updatedAt: Date;
  private deletedAt?: Date;

  constructor(createdAt: Date, updatedAt: Date, deletedAt?: Date) {
    if (updatedAt < createdAt) {
      throw ErrInvalidTimestamps;
    }

    if (deletedAt) {
      if (deletedAt < createdAt) {
        throw ErrInvalidTimestamps;
      }

      if (deletedAt < updatedAt) {
        throw ErrInvalidTimestamps;
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
