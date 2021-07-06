export interface Channel {
  id: string;
  name: string;
  keywords: string[];
  owner: Owner;
}

export interface Owner {
  id: number;
  username: string;
}

export type AccessToken = string;
export interface LoginResponseErrors {
  username?: string;
  password?: string;
}
export interface LoginResponse {
  accessToken?: AccessToken;
  errors?: LoginResponseErrors;
}

export interface MeResponseErrors {
  uid: string;
}
export interface MeResponse {
  username?: string;
  errors?: MeResponseErrors;
}
