const data = {
  latitude: 53.5458874,
  longitude: -113.5034304,
  timezone: "America/Edmonton",
  currently: {
    summary: "Light Snow",
    temperature: 20.33,
  },
};

const {
  currently: {summary, temperature},
} = data;

console.info(`${summary} and ${temperature}ºF`);
