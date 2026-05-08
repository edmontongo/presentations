async function fetchHumans(url) {
  const response = await fetch(url);
  return response.text();
}

const express = require("express");
const app = express();
const port = 3000;

const urls = [
  "https://www.google.com/humans.txt",
  "https://www.netflix.com/humans.txt",
  "https://medium.com/humans.txt",
];

app.get("/", async (req, res) => {
  const promises = urls.map((url) => fetchHumans(url));
  const data = await Promise.all(promises);
  res.setHeader("content-type", "text/plain");
  res.send(data.join("\n\n"));
});

app.listen(port, () => {
  console.info(`Listening on ${port}`);
});
