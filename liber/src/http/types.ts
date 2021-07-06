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
export interface TokenResponseErrors {
  username?: string;
  password?: string;
}
export interface TokenResponse {
  accessToken?: AccessToken;
  errors?: TokenResponseErrors;
}
