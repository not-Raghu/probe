import express from "express";
import http from "http";
import { WebSocketServer } from "ws";

const app = express();

const server = http.createServer(app);

const wss = new WebSocketServer({ server });

wss.on("connection", (ws) => {
  console.log("Client connected");

  ws.on("error", () => {
    console.log("error on ws");
  });

  ws.on("ping", () => {
    console.log("pinged bro");
  });
}).on('close',()=>{
    console.log("ws server closed")
});

app.get('/',)
// Start the server
server.listen(3000, () => {
  console.log("Server running on port 3000");
});
