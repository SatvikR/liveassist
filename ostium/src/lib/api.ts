import { HTTPClient } from "@liveassist/liber";

const apiUrl = process.env.API_URL;

export const api = !!apiUrl ? new HTTPClient(apiUrl) : new HTTPClient();
