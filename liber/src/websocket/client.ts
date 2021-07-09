import WebSocket from "isomorphic-ws";

const BASE_URL = "wss://lapi.satvikreddy.com";

export class MessagingClient {
  private url: string;
  private socket: WebSocket;
  private callback: (message: string) => any;

  public constructor(callback: (message: string) => any, url?: string) {
    this.url = url || BASE_URL;
    this.callback = callback;
  }

  public connect() {
    this.socket = new WebSocket(this.url);
    this.socket.onmessage = (e) => this.callback(e.data as string);
  }

  public send(message: string) {
    this.socket.send(message);
  }
}
