const amqp = require("amqplib");

(async () => {
  const RABBIT_URL = "amqp://admin:admin@rabbitmq:5672";
  const connection = await amqp.connect(RABBIT_URL);
  const channel = await connection.createChannel();
  await channel.assertExchange("events-topic-exchange", "topic", { durable: true });

  let i = 0;
  while (true) {
    const msg = { source: "producer-node", message: `event ${i++}` };
    channel.publish("events-topic-exchange", "events.node", Buffer.from(JSON.stringify(msg)), {
      persistent: true,
    });
    console.log("Sent:", msg);
    await new Promise((r) => setTimeout(r, 5000));
  }
})();
