import { Error } from '../errors';
import { Metadata } from '../collections';

const rePart = new RegExp('^[a-z]+[a-z0-9_]*[a-z0-9]+$', 'i');

export const ErrInvalidPath = Error.define('path.invalid');

function isPartValid(part: string, separator: string, ...wildcards: string[]): boolean {
  if (part === separator) {
    return true;
  }

  if (wildcards.some((wildcard) => part === wildcard)) {
    return true;
  }

  return rePart.test(part);
}

export class Path {
  private parts: string[];
  private separator: string;
  private wildcards: string[];

  constructor(path: string, separator: string, ...wildcards: string[]) {
    const parts = path.split(separator);

    if (parts.length == 0) {
      throw Error.create(ErrInvalidPath, 'pfull path is empty');
    }

    this.parts = parts.map((part) => {
      part = part.toLowerCase();

      if (!isPartValid(part, separator, ...wildcards)) {
        throw Error.create(
          ErrInvalidPath,
          'path part has invalid characters',
          Metadata.with('path', path)
            .and('parts', parts)
            .and('separator', separator)
            .and('wildcards', wildcards),
        );
      }

      return part;
    });

    this.separator = separator;
    this.wildcards = wildcards;
  }

  getParts(): string[] {
    return this.parts;
  }

  getSeparator(): string {
    return this.separator;
  }

  getWildcards(): string[] {
    return this.wildcards;
  }

  toString(): string {
    return this.parts.join(this.separator);
  }

  equals(other: Path): boolean {
    if (this.parts.length != other.parts.length) {
      return false;
    }

    return this.parts.every((p1, i) => {
      const p2 = other.parts[i];

      return p1 === p2;
    });
  }
}
