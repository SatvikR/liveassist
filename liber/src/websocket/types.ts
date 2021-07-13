export interface Message {
  id: string;
  createdAt: string;
  text: string;
  owner: MessageOwner;
  channelId: string;
}

export interface MessageOwner {
  id: number;
  username: string;
}
