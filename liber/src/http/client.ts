import axios, { AxiosInstance } from "axios";
import { ChannelService } from "./channels";

const BASE_URL = "https://lapi.satvikreddy.com";

export class HTTPClient {
  private api: AxiosInstance;
  private _channels: ChannelService;

  public constructor(url?: string) {
    this.api = axios.create({
      baseURL: url || BASE_URL,
    });
    this._channels = new ChannelService(this.api);
  }

  get channels(): ChannelService {
    return this._channels;
  }
}
