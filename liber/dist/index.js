"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.HTTPClient = void 0;
const axios_1 = __importDefault(require("axios"));
const channels_1 = require("./channels");
const BASE_URL = "https://lapi.satvikreddy.com";
class HTTPClient {
    constructor(url) {
        this.api = axios_1.default.create({
            baseURL: url || BASE_URL,
        });
        this._channels = new channels_1.ChannelService(this.api);
    }
    get channels() {
        return this._channels;
    }
}
exports.HTTPClient = HTTPClient;
