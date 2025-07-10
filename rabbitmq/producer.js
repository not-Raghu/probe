import ampqp from "amqplib";

async function main() {
  const conn = await ampqp.connect("amqp://localhost:5672");
  const channel = await conn.createChannel();
  await channel.assertQueue("num");

  for (let i = 0; i < 10; i++) {
    const eventSuccess = channel.sendToQueue("num", Buffer.from(i.toString()));
    if (!eventSuccess) {
      console.log("error");
    }
  }
  console.log("done");
}

main();
