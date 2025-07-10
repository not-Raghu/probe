import amqplib from "amqplib";

async function main() {
  const conn = await amqplib.connect("amqp://localhost:5672");

  const channel = await conn.createChannel();
  const data = await channel.consume("num", (message) => {
    const input = message.content.toString();
    console.log(input);

  });
  console.log(data);
}

main();
