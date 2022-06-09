import { Error } from '../errors';
import { Metadata } from '../collections';

export const ErrInvalidVersion = Error.define('version.invalid');

export class Version {
  private version: number;
  private updated: boolean;

  constructor(version: number) {
    if (version < 1) {
      throw Error.create(
        ErrInvalidVersion,
        'version is smaller than 1',
        Metadata.with('version', version),
      );
    }

    this.version = version;
    this.updated = false;
  }

  static initVersion(): Version {
    const version = new Version(1);
    version.updated = true;
    return version;
  }

  value(): number {
    return this.version;
  }

  equals(other: Version): boolean {
    return this.version === other.version;
  }

  incr(): Version {
    if (this.updated) {
      return this;
    }

    const version = new Version(this.version + 1);
    version.updated = true;

    return version;
  }
}
