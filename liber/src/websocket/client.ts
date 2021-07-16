import WebSocket from "isomorphic-ws";
import { Message } from "./types";

const BASE_URL = "wss://lapi.satvikreddy.com/messages/ws/";

export class MessagingClient {
  private url: string;
  private socket: WebSocket | undefined;
  private messageCallback: (message: Message) => any;
  private connectCallback: () => any;
  private messagesCallback: (messages: Message[]) => any;

  public readonly PAGE_SIZE = 30;

  public constructor(
    messageCallback: (message: Message) => any,
    connectCallback: () => any,
    messagesCallback: (messages: Message[]) => any,
    url?: string
  ) {
    this.url = url || BASE_URL;
    if (this.url[this.url.length - 1] != "/") {
      this.url = this.url + "/";
    }
    this.messageCallback = messageCallback;
    this.connectCallback = connectCallback;
    this.messagesCallback = messagesCallback;
  }

  public handleMessage(e: WebSocket.MessageEvent) {
    const data = JSON.parse(e.data as string);
    if (!data) {
      this.messagesCallback([]);
      return;
    }

    if (Array.isArray(data)) {
      this.messagesCallback(data);
      return;
    }

    this.messageCallback(data);
  }

  public connect(channel: string, token: string) {
    this.socket = new WebSocket(
      `${this.url}?channel=${channel}&token=${token}`
    );
    this.socket.onmessage = (e) => this.handleMessage(e);
    this.socket.onopen = (_e) => this.connectCallback();
  }

  public send(text: string) {
    this.socket.send(
      JSON.stringify({
        text,
      })
    );
  }

  public close() {
    this.socket.close();
  }

  public loadMore(lastMessage: Message) {
    this.socket.send(
      JSON.stringify({
        loadMore: lastMessage.createdAt,
      })
    );
  }
}
