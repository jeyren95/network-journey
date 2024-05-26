import { type IpHopsOptions } from "../types";

const API_ENDPOINT = "http://localhost:8080/ip";

export const getGeolocationHops = (body: IpHopsOptions, signal?: AbortSignal) =>
  fetch(API_ENDPOINT, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      accept: "application/json",
    },
    body: JSON.stringify(body),
    signal,
  });
