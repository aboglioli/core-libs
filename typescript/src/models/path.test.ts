import { Path } from './path';

describe('Path', () => {
  test('equals', () => {
    expect(new Path('one.two.three', '.').equals(new Path('one.two.three', '.'))).toBe(
      true,
    );
    expect(new Path('one.two.three', '.').equals(new Path('one.three.two', '.'))).toBe(
      false,
    );
  });
});
