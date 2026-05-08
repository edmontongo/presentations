async function fetchHumans(url) {
  const response = await fetch(url);
  return response.text();
}

const url = "https://www.google.com/humans.txt";

const express = require("express");
const app = express();
const port = 3000;

app.get("/", async (req, res) => {
  const data = await fetchHumans(url);
  res.send(data);
});

app.listen(port, () => {
  console.info(`Listening on ${port}`);
});
