import { BaseService } from "./service";
import { Channel } from "./types";

export class ChannelService extends BaseService {
  private readonly BASE_PATH = "/api/channels";

  public async list(): Promise<Channel[]> {
    const res = await this.api.get<Channel[]>(`${this.BASE_PATH}/`);
    return res.data;
  }
}
