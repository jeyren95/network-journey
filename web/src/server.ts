import {
  createServer,
  type ServerResponse,
  type OutgoingHttpHeaders,
  type IncomingMessage,
} from "http";
import { join } from "path";
import { readFileSync } from "fs";

// read files
const BUFFERS = {
	homePage: readFileSync(join(__dirname, "..", "src", "pages", "index.html")),
	notFoundPage: readFileSync(join(__dirname, "..", "src", "pages", "404.html")),
	stylesheet: readFileSync(join(__dirname, "..", "src", "index.css")),
	leafletScript: readFileSync(join(__dirname, "leaflet.js")),
}

enum URL {
	HOME_PAGE = "/",
	STYLESHEET = "/src/index.css",
	LEAFLET_SCRIPT = "/dist/leaflet.js",
}

enum MimeType {
	JS = "text/javascript",
	CSS = "text/css",
	HTML = "text/html",
}

const writeStream = (
  res: ServerResponse<IncomingMessage>,
  headers: OutgoingHttpHeaders,
  statusCode: number,
  buffer: Buffer,
) => {
  res.writeHead(statusCode, headers);
  res.write(buffer);
  res.end();
};

const server = createServer((req, res) => {
  const url = req.url;
  let headers: OutgoingHttpHeaders = {};

  switch (url) {
    case URL.LEAFLET_SCRIPT:
      headers["Content-Type"] = MimeType.JS;
      writeStream(res, headers, 200, BUFFERS.leafletScript);
      break;
    case URL.STYLESHEET:
      headers["Content-Type"] = MimeType.CSS;
      writeStream(res, headers, 200, BUFFERS.stylesheet);
      break;
		case URL.HOME_PAGE:
      headers["Content-Type"] = MimeType.HTML;
			writeStream(res, headers, 200, BUFFERS.homePage);
			break;
    default:
      headers["Content-Type"] = MimeType.HTML;
  		writeStream(res, headers, 404, BUFFERS.notFoundPage);
	}
});

const PORT = 3000;
const HOST = "127.0.0.1";

server.listen(PORT, HOST, () => {
  console.log("Server listening on port 3000");
});