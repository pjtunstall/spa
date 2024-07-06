const http = require("http");
const fs = require("fs");
const path = require("path");

const server = http.createServer((req, res) => {
  let filePath = req.url === "/" ? "index.html" : req.url;
  filePath = filePath.replace(/^\//, ""); // Remove leading slash for security.
  let safeFilePath = path.normalize(filePath).replace(/^(\.\.[\/\\])+/, "");
  safeFilePath = path.join(__dirname, safeFilePath);

  if (!safeFilePath.startsWith(__dirname)) {
    res.writeHead(400);
    res.end("Invalid path");
    return;
  }

  fs.access(safeFilePath, fs.constants.F_OK, (err) => {
    if (err) {
      // File does not exist, serve index.html.
      safeFilePath = path.join(__dirname, "index.html");
    }

    // Determine content type.
    let contentType = "text/html"; // Default content type.
    const ext = path.extname(safeFilePath);
    if (ext === ".js") {
      contentType = "text/javascript";
    } else if (ext === ".css") {
      contentType = "text/css";
    }

    fs.readFile(safeFilePath, (err, content) => {
      if (err) {
        res.writeHead(500);
        res.end(`Server Error: ${err.code}`);
        return;
      }
      res.writeHead(200, { "Content-Type": contentType });
      res.end(content, "utf-8");
    });
  });
});

const port = 3000;
server.listen(port, () => {
  console.log(`Server running at http://localhost:${port}/`);
});
