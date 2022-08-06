import { Error, ErrorCode } from './error';
import { Metadata } from './metadata';

describe('Error', () => {
  test('create new error', () => {
    const code = new ErrorCode('custom_code');
    const err = Error.create(code, 'custom error');

    expect(err.getCode()).toBe(code);
    expect(err.getMessage()).toBe('custom error');
    expect(err.getCause()).toBeUndefined();
  });

  test('with metadata', () => {
    const code = new ErrorCode('custom_code');
    const err = Error.create(
      code,
      'custom error',
      Metadata.with('key1', 'value1').and('key2', 2),
    );

    expect(err.getMetadata().values()).toEqual({
      key1: 'value1',
      key2: 2,
    });
  });

  test('wrap error', () => {
    const innerErr = Error.create(new ErrorCode('inner_err'), 'inner error');

    const err = Error.wrap(new ErrorCode('err'), innerErr, 'error');

    expect(err.getCause()).toBe(innerErr);
  });
});
