const amqp = require("amqplib");

(async () => {
  const conn = await amqp.connect("amqp://admin:admin@rabbitmq:5672");
  const ch = await conn.createChannel();
  await ch.assertExchange("events-exchange", "fanout", { durable: true });

  const q = await ch.assertQueue("events-node", { durable: true });
  await ch.bindQueue(q.queue, "events-exchange", "");

  console.log(" [*] Waiting for messages...");
  ch.consume(q.queue, (msg) => {
    if (msg) {
      const data = JSON.parse(msg.content.toString());
      console.log(`Received from ${data.source}: ${data.message}`);
      ch.ack(msg);
    }
  });
})();

