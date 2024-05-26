import {
  LatLng,
  Map,
  Marker,
  TileLayer,
  Polyline,
  type PolylineOptions,
  type TileLayerOptions,
  type LatLngExpression,
} from "https://cdn.jsdelivr.net/npm/leaflet@1.9.4/dist/leaflet-src.esm.js";
import type { GetIpHopsRes, IpHopsOptions } from "./types";
import { getIpHops } from "./services";

// initialize map
const TILE_LAYER_URL = "https://tile.openstreetmap.org/{z}/{x}/{y}.png";
const DEFAULT_MAX_ZOOM = 19;
const ATTRIBUTION =
  '&copy; <a href="http://www.openstreetmap.org/copyright">OpenStreetMap</a>';

const map = new Map("map");
const tileLayerOptions: TileLayerOptions = {
  maxZoom: DEFAULT_MAX_ZOOM,
  attribution: ATTRIBUTION,
  noWrap: true,
};
const tileLayer = new TileLayer(TILE_LAYER_URL, tileLayerOptions);
tileLayer.addTo(map);
map.fitWorld();

// add ui layer methods
const MARKER_ZOOM_AROUND = 15;
const addMarker = (latlng: LatLngExpression) => {
  const marker = new Marker(latlng);
  marker.on("click", () => {
    map.setZoomAround(latlng, MARKER_ZOOM_AROUND);
  });
  marker.addTo(map);
};

// add vector layer methods
const VIEW_ZOOM = 10;
const addPolyline = (latlngs: LatLngExpression[], options: PolylineOptions) => {
  const polyline = new Polyline(latlngs, options);
  polyline.addTo(map);
  map.setView(latlngs[0], VIEW_ZOOM);
};

// drawing the map
const form = document.querySelector("form");
(form as HTMLFormElement).addEventListener("submit", (event) => {
  event.preventDefault();

  const hostname = (<HTMLInputElement>document.querySelector("#hostname"))
    .value;
  const maxHops = (<HTMLInputElement>document.querySelector("#max-hops")).value;
  const waitTime = (<HTMLInputElement>document.querySelector("#wait-time"))
    .value;
  const requestBody: IpHopsOptions = {
    hostname,
    waitTime: Number(waitTime),
    maxHops: Number(maxHops),
  };

  getIpHops(requestBody)
    .then((res) => res.json())
    .then((data: GetIpHopsRes) => {
      const latlngs = data.geolocations.map((h) => new LatLng(h.lat, h.lon));
      latlngs.forEach((l: LatLng) => addMarker(l));
      addPolyline(latlngs, { color: "red" });
    })
    .catch((err) => console.log(err));
});
