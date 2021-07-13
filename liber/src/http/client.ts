import axios, { AxiosInstance } from "axios";
import { ChannelService } from "./channels";
import { TokenService } from "./token";
import { UserService } from "./user";

const BASE_URL = "https://lapi.satvikreddy.com";

export class HTTPClient {
  private api: AxiosInstance;
  private _channels: ChannelService;
  private _users: UserService;
  private _tokens: TokenService;

  public constructor(url?: string) {
    this.api = axios.create({
      baseURL: url || BASE_URL,
      withCredentials: true,
    });
    this._channels = new ChannelService(this.api);
    this._users = new UserService(this.api);
    this._tokens = new TokenService(this.api);
  }

  get channels(): ChannelService {
    return this._channels;
  }

  get users(): UserService {
    return this._users;
  }

  get tokens(): TokenService {
    return this._tokens;
  }
}
