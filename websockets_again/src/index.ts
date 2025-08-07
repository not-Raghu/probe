import express from "express";
import http from "http";
import { matchesGlob } from "path";
import { WebSocketServer } from "ws";

const app = express();

const server = http.createServer(app);

const wss = new WebSocketServer({server:server})

wss.on("connection",(ws)=>{
    console.log("client connected")
    ws.send("hi");
    ws.addEventListener("message",(message)=>{
        console.log(message.data)
    })  
    ws.addEventListener('close',(message)=>{
        console.log(message.reason)
    })
    setTimeout(()=>{
        ws.send("sending after 2 seconds")
    },1000)
})


server.listen(8000,()=>{
    console.log("server listening on port 8000")
})
