import { Error as CustomError } from '../errors';

type Ok<T> = { ok: true; value: T };
type Err<E> = { ok: false; err: E };

type ResultValue<T, E> = Ok<T> | Err<E>;

export class Result<T, E = CustomError> {
  private constructor(private readonly v: ResultValue<T, E>) {}

  static ok<T, E>(value: T): Result<T, E> {
    return new Result<T, E>({ ok: true, value });
  }

  static err<T, E>(err: E): Result<T, E> {
    return new Result<T, E>({ ok: false, err });
  }

  static fromThrowable<T, E>(fn: () => T, mapErr?: (err: E) => E): Result<T, E> {
    try {
      return Result.ok(fn());
    } catch (err) {
      return Result.err(mapErr ? mapErr(err as E) : (err as E));
    }
  }

  isOk(): boolean {
    return this.v.ok;
  }
  isErr(): boolean {
    return !this.v.ok;
  }

  unwrap(): T {
    if (this.v.ok) {
      return this.v.value;
    }

    throw this.v.err;
  }

  unwrapOrElse(defaultValue: T): T {
    if (this.v.ok) {
      return this.v.value;
    }

    return defaultValue;
  }

  error(): E | undefined {
    if (this.v.ok) {
      return undefined;
      // throw ErrNotError;
    }

    return this.v.err;
  }

  map<U>(fn: (v: T) => U): Result<U, E> {
    if (this.v.ok) {
      return Result.ok(fn(this.v.value));
    }

    return Result.err(this.v.err);
  }

  mapErr<F>(fn: (err: E) => F): Result<T, F> {
    if (this.v.ok) {
      return Result.ok(this.v.value);
    }

    return Result.err(fn(this.v.err));
  }
}
