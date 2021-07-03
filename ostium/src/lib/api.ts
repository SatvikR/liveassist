import { HTTPClient } from "@liveassist/liber";

const apiUrl = process.env.NEXT_PUBLIC_API_URL;

export const api = !!apiUrl ? new HTTPClient(apiUrl) : new HTTPClient();
