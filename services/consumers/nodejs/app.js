const amqp = require("amqplib");

(async () => {
  const RABBIT_URL = "amqp://admin:admin@rabbitmq:5672";
  const connection = await amqp.connect(RABBIT_URL);
  const channel = await connection.createChannel();

  await channel.assertExchange("events-topic-exchange", "topic", { durable: true });
  const q = await channel.assertQueue("events-node", { durable: true });
  await channel.bindQueue(q.queue, "events-topic-exchange", "events.node");

  console.log(" [*] Waiting for Node events...");
  channel.consume(q.queue, (msg) => {
    const data = JSON.parse(msg.content.toString());
    console.log(`[Node consumer] Received from ${data.source}: ${data.message}`);
    channel.ack(msg);
  });
})();
