export class AccessToken {
  private static token: string | null = null;

  static get value(): string | null {
    return this.token;
  }

  static set value(ntoken: string) {
    this.token = ntoken;
  }
}
