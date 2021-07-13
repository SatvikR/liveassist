import { HTTPClient } from "@liveassist/liber";

const apiUrl = process.env.NEXT_PUBLIC_API_URL;
export const MSG_URL = process.env.NEXT_PUBLIC_WS_URL || null;

export const api = !!apiUrl ? new HTTPClient(apiUrl) : new HTTPClient();
