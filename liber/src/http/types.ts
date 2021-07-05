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
export interface TokenResponse {
  accessToken: string;
}
