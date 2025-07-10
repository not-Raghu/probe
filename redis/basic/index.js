import axios from "axios";
import express, { json } from "express";
import { Redis } from "ioredis";

const app = express();
const client = new Redis();
app.get("/api/posts", async (req, res) => {
  const isCached = await client.get("posts");
  if (isCached) {
    return res.json(JSON.parse(isCached));
  }

  const response = await axios.get(
    "https://jsonplaceholder.typicode.com/posts"
  );
  const data = response.data;
  await client.set("posts", JSON.stringify(data));
  await client.expire("posts", 10 );
  return res.json(data);
});

app.listen(3000, () => {
  console.log("server listening");
});
