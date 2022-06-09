import { Metadata } from '../collections';

export class ErrorCode {
  private code: string;

  constructor(code: string) {
    this.code = code;
  }

  toString(): string {
    return this.code;
  }
}

export class Error {
  private code: ErrorCode;
  private message: string;
  private cause?: unknown;
  private metadata: Metadata;

  constructor(
    code: ErrorCode,
    message: string,
    cause?: unknown,
    ...metadata: Metadata[]
  ) {
    let m = new Metadata();
    for (const meta of metadata) {
      m = m.merge(meta);
    }

    this.code = code;
    this.message = message;
    this.cause = cause;
    this.metadata = m;
  }

  static define(code: string): ErrorCode {
    return new ErrorCode(code);
  }

  static create(code: ErrorCode, message: string, ...metadata: Metadata[]): Error {
    return new Error(code, message, undefined, ...metadata);
  }

  static wrap(
    code: ErrorCode,
    cause: unknown,
    message: string,
    ...metadata: Metadata[]
  ): Error {
    return new Error(code, message, cause, ...metadata);
  }

  getCode(): ErrorCode {
    return this.code;
  }

  getMessage(): string {
    return this.message;
  }

  getCause(): unknown | undefined {
    return this.cause;
  }

  getMetadata(): Metadata {
    return this.metadata;
  }

  isAny(...codes: ErrorCode[]): boolean {
    return codes.some((code) => this.code === code);
  }

  toString(): string {
    return `${this.code.toString()}: ${this.message}`;
  }
}
