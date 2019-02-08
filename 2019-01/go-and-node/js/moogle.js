const axios = require("axios");
const url = "https://www.moogle.com/humans.txt";

async function fetchHumans(url) {
  const response = await axios.get(url);
  return response.data;
}

const express = require("express");
const app = express();
const port = 3000;

app.get("/", async (req, res) => {
  try {
    const data = await fetchHumans(url);
    res.send(data);
  } catch (err) {
    console.error(err.message);
    res.status(502).send("Bad gateway");
  }
});

app.listen(port, () => {
  console.info(`Listening on ${port}`);
});
