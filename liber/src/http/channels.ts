import { AxiosError } from "axios";
import { BaseService } from "./service";
import { Channel, CreateChannelResponse } from "./types";

export class ChannelService extends BaseService {
  private readonly BASE_PATH = "/api/channels";

  public async list(): Promise<Channel[]> {
    const res = await this.api.get<Channel[]>(`${this.BASE_PATH}/`);
    return res.data;
  }

  public async get(id: string): Promise<Channel> {
    try {
      const res = await this.api.get<Channel>(`${this.BASE_PATH}/${id}`);
      return res.data;
    } catch (_e) {
      const error = _e as AxiosError<Channel>;
      return error.response.data;
    }
  }

  public async create(
    name: string,
    keywords: string[],
    token: string
  ): Promise<CreateChannelResponse> {
    try {
      const res = await this.api.post<CreateChannelResponse>(
        `${this.BASE_PATH}/`,
        {
          name,
          keywords,
        },
        {
          headers: {
            Authorization: `Bearer ${token}`,
          },
        }
      );
      return res.data;
    } catch (_e) {
      const error = _e as AxiosError<CreateChannelResponse>;
      return error.response.data;
    }
  }
}
