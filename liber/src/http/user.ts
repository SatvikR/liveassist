import { TokenResponse } from "./types";
import { BaseService } from "./service";
import { AxiosError } from "axios";

export class UserService extends BaseService {
  private readonly BASE_PATH = "/api/users";

  public async login(
    username: string,
    password: string
  ): Promise<TokenResponse> {
    try {
      const res = await this.api.post<TokenResponse>(
        `${this.BASE_PATH}/login`,
        {
          username,
          password,
        }
      );
      return res.data;
    } catch (_e) {
      const error = _e as AxiosError<TokenResponse>;
      console.log(error.response.data);
    }
  }
}
