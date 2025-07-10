import express from "express";
import { PrismaClient } from "./generated/prisma";

const prisma = new PrismaClient()



const app = express();

// app.use(express.json());

// app.post("/post", async (req, res) => {
//   const image = prisma.images.create({
//     data: {
//       url: "https://res.cloudinary.com/delcjqr59/image/upload/v1745124623/qwgcaxbtkqwa6sqsuxjm.jpg",
//       publicId: "qwgcaxbtkqwa6sqsuxjm",
//     },
//   });
//   res.json({
//     image
//   })
// });


// app.listen(3000)
async   function main(){
  await prisma.$connect();
  console.log("connected")
}


main()