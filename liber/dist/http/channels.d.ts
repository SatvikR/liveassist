import { AxiosInstance } from "axios";
export declare class ChannelService {
    private readonly BASE_PATH;
    private api;
    constructor(api: AxiosInstance);
    list(): Promise<Channel[]>;
}
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
