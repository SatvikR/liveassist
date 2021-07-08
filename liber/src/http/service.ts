import { AxiosInstance } from "axios";

export class BaseService {
  private _api: AxiosInstance;

  public constructor(api: AxiosInstance) {
    this._api = api;
  }

  protected get api(): AxiosInstance {
    return this._api;
  }
}
