import { AxiosError } from "axios";
import { BaseService } from "./service";
import { LoginResponse, MeResponse, SignupResponse } from "./types";

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

  public async signup(
    username: string,
    email: string,
    password: string
  ): Promise<SignupResponse> {
    try {
      const res = await this.api.post<SignupResponse>(
        `${this.BASE_PATH}/signup`,
        {
          username,
          email,
          password,
        }
      );
      return res.data;
    } catch (_e) {
      const error = _e as AxiosError<SignupResponse>;
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
