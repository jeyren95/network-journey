export type IpHopsOptions = {
  hostname: string;
  waitTime: number;
  maxHops: number;
};

export type IpHop = {
  hostname: string;
  ip: string;
  isIpPrivate: boolean;
  returnTime: string;
};

export type Geolocation = {
  status: string;
  message: string;
  country: string;
  countryCode: string;
  region: string;
  regionName: string;
  city: string;
  zip: string;
  lat: number;
  lon: number;
  timezone: number;
  isp: string;
  org: string;
  as: string;
};

export type GetIpHopsRes = {
  geolocations: Geolocation[];
  ipHops: IpHop[];
};
