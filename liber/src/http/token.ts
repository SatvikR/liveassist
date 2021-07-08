import { BaseService } from "./service";
import { RefreshResponse } from "./types";

export class TokenService extends BaseService {
  private readonly BASE_PATH = "/api/tokens";

  public async refresh(): Promise<string | null> {
    try {
      const res = await this.api.put<RefreshResponse>("/api/tokens/refresh");
      return res.data.accessToken;
    } catch (_e) {
      return null;
    }
  }
}
