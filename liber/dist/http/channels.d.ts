import { AxiosInstance } from "axios";
import { Channel } from "./types";
export declare class ChannelService {
    private readonly BASE_PATH;
    private api;
    constructor(api: AxiosInstance);
    list(): Promise<Channel[]>;
}
