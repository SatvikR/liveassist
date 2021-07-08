import { AxiosError } from "axios";
import { BaseService } from "./service";
import { Channel, CreateChannelResponse } from "./types";

export class ChannelService extends BaseService {
  private readonly BASE_PATH = "/api/channels";

  public async list(): Promise<Channel[]> {
    const res = await this.api.get<Channel[]>(`${this.BASE_PATH}/`);
    return res.data;
  }

  public async create(
    name: string,
    keywords: string[]
  ): Promise<CreateChannelResponse> {
    try {
      const res = await this.api.post<CreateChannelResponse>(
        `${this.BASE_PATH}/`,
        {
          name,
          keywords,
        }
      );
      return res.data;
    } catch (_e) {
      const error = _e as AxiosError<CreateChannelResponse>;
      return error.response.data;
    }
  }
}
