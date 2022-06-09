export interface Map {
  [k: string]: unknown;
}

export class Metadata {
  private metadata: Map;

  constructor() {
    this.metadata = {};
  }

  static with(key: string, value: unknown): Metadata {
    return new Metadata().and(key, value);
  }

  and(key: string, value: unknown): Metadata {
    this.metadata[key] = value;
    return this;
  }

  merge(other: Metadata): Metadata {
    this.metadata = {
      ...this.metadata,
      ...other.metadata,
    };

    return this;
  }

  values(): Map {
    return this.metadata;
  }
}
