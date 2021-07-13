export class QueryKeys {
  static get channels(): string {
    return "channels";
  }

  static get me(): string {
    return "me";
  }

  static channel(id: string): string[] {
    return ["channel", id];
  }
}
