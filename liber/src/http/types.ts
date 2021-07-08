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

export interface SignupResopnseErrors {
  username?: string;
  email?: string;
  password?: string;
}
export interface SignupResponse {
  accessToken?: AccessToken;
  errors?: SignupResopnseErrors;
}
export interface MeResponseErrors {
  uid: string;
}
export interface MeResponse {
  username?: string;
  errors?: MeResponseErrors;
}

export interface RefreshResponse {
  accessToken?: string;
  error?: string;
}

export interface CreateChannelResponse {
  id?: string;
  error?: string;
}
