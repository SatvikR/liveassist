import jwtDecode, { JwtPayload } from "jwt-decode";

export class AccessToken {
  private token: string | null;
  private exp: number;
  private static instance: AccessToken;

  private constructor() {
    this.token = null;
  }

  public static getInstance(): AccessToken {
    if (!AccessToken.instance) {
      AccessToken.instance = new AccessToken();
    }
    return AccessToken.instance;
  }

  public get value(): string | null {
    return this.token;
  }

  public set value(ntoken: string) {
    this.token = ntoken;
    this.exp = jwtDecode<JwtPayload>(this.token).exp;
  }

  public reset() {
    this.token = null;
  }

  public isExp(): boolean {
    if (this.token == null) {
      return true;
    }
    return this.exp < Date.now() / 1000;
  }
}
