import Redis from "ioredis";
import { config } from "dotenv";
import express from "express";
config();

const app = express();
const client = new Redis(process.env.REDIS_URL);

app.post("/put", (_, res) => {
  const data = redisPut();
  return res.json({
    data,
  });
});

app.get("/delete", (_, res) => {
  const data = redisDelete();
  return res.json({
    data,
  });
});


async function redisPut() {
  const data = await client.set("data", 1);
  return data;
}

async function redisDelete() {
  const data = await client.del("data");
  return data;
}

app.listen(3000, () => {
  console.log("server running on port 3000");
});
