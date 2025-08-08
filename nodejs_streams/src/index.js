import express from "express";
import monitor from "express-status-monitor";
import fs from "fs";
import cors from 'cors';


const app = express();
app.use(monitor());
app.use(cors())

app.get("/", (_, res) => {
    console.log("user came")
  fs.readFile("/Users/raghu/Downloads/dump/bigfile.txt", (err, data) => {
    if (err) {
      console.log(err);
    }
    res.(data);
  });
});

app.listen(3000, () => {
  console.log("listening on port 3000");
});
