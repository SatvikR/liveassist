import { ChannelService } from "./channels";
export declare class HTTPClient {
    private api;
    private _channels;
    constructor(url?: string);
    get channels(): ChannelService;
}
