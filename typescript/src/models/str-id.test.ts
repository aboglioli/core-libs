import { StrId } from './str-id';

describe('StrId', () => {
  test('generate slug', () => {
    const slug = StrId.generateSlug('Slug Example!');
    expect(slug.value()).toBe('slug-example');

    expect(() => StrId.generateSlug('')).toThrow();
  });
});
