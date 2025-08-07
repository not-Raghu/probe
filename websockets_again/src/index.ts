// import express from "express";
// import http from "http";
// import { Event, WebSocketEventMap, WebSocketServer } from "ws";
// const app = express();

// const server = http.createServer(app);

// const wss = new WebSocketServer({ server: server });

// // function sendMessage(ws:any) {
// //     let count = 0;
// //     const interval = setInterval(() => {
// //         if (ws.readyState === ws.OPEN) {
// //             ws.send(`${count}`);
// //             count++;
// //         } else {
// //             clearInterval(interval);
// //         }
// //     }, 1000);
// // }



// wss.on("connection",(ws)=>{
    
//     ws.send(wss.clients.size)

//     ws.on('close',()=>{
//         console.log(wss.clients.size);
//     })
// })


// server.listen(8000, () => {
//   console.log("server listeing on port 8000");
// });


import express from "express";
import http from "http";
import { WebSocketServer } from "ws";

const app = express();
const server = http.createServer(app);

const wss = new WebSocketServer({ server });

// Broadcast helper
function broadcastCount() {
    const count = wss.clients.size;
    for (const client of wss.clients) {
        if (client.readyState === client.OPEN) {
            client.send(JSON.stringify({ count }));
        }
    }
}

wss.on("connection", (ws) => {
    console.log("Client connected");
    broadcastCount(); 

    ws.on("close", () => {
        console.log("Client disconnected");
        broadcastCount(); 
    });
});

server.listen(8000, () => {
    console.log("Server listening on port 8000");
});
