import { Result } from './result';

describe('Result', () => {
  test('ok', () => {
    const res = Result.ok('hello');

    expect(res.unwrap()).toBe('hello');
    expect(res.isOk()).toBeTruthy();
    expect(res.error()).toBeUndefined();
    expect(res.isErr()).toBeFalsy();
  });

  test('error', () => {
    const res = Result.err('hello');

    expect(() => res.unwrap()).toThrow();
    expect(res.isOk()).toBeFalsy();
    expect(res.error()).toBe('hello');
    expect(res.isErr()).toBeTruthy();
    expect(res.unwrapOrElse(123)).toBe(123);
  });

  test('map ok and err', () => {
    const makeSomething = (ok: boolean): Result<string, Error> => {
      if (ok) {
        return Result.ok('Hello');
      }

      return Result.err(new Error('Error'));
    };

    let res = makeSomething(true)
      .map((str) => str?.toUpperCase())
      .map((str) => `OK: ${str}`)
      .mapErr((err) => new Error(err?.message.toLowerCase()))
      .mapErr((err) => new Error(`ERR: ${err?.message}`));

    expect(res.unwrap()).toBe('OK: HELLO');
    expect(res.error()).toBeUndefined();

    res = makeSomething(false)
      .map((str) => str?.toUpperCase())
      .map((str) => `OK: ${str}`)
      .mapErr((err) => new Error(err?.message.toLowerCase()))
      .mapErr((err) => new Error(`ERR: ${err?.message}`));

    expect(() => res.unwrap()).toThrow();
    expect(res.error()).toEqual(new Error('ERR: error'));
  });
});
