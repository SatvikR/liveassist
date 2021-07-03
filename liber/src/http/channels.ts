import { AxiosInstance } from "axios";
import { Channel } from "./types";

export class ChannelService {
  private readonly BASE_PATH = "/api/channels/";
  private api: AxiosInstance;

  public constructor(api: AxiosInstance) {
    this.api = api;
  }

  public async list(): Promise<Channel[]> {
    const res = await this.api.get<Channel[]>(`${this.BASE_PATH}`);
    return res.data;
  }
}
