import { LoginResponse, MeResponse } from "./types";
import { BaseService } from "./service";
import { AxiosError } from "axios";

export class UserService extends BaseService {
  private readonly BASE_PATH = "/api/users";

  public async login(
    username: string,
    password: string
  ): Promise<LoginResponse> {
    try {
      const res = await this.api.post<LoginResponse>(
        `${this.BASE_PATH}/login`,
        {
          username,
          password,
        }
      );
      return res.data;
    } catch (_e) {
      const error = _e as AxiosError<LoginResponse>;
      return error.response.data;
    }
  }

  public async me(token: string): Promise<MeResponse> {
    try {
      const res = await this.api.get<MeResponse>(`${this.BASE_PATH}/me`, {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      });
      return res.data;
    } catch (_e) {
      const error = _e as AxiosError<MeResponse>;
      return error.response.data;
    }
  }
}
