const amqp = require("amqplib");

(async () => {
  try {
    const RABBIT_URL = "amqp://admin:admin@rabbitmq:5672";
    const connection = await amqp.connect(RABBIT_URL);
    const channel = await connection.createChannel();

    // Declarar exchange fanout y durable
    await channel.assertExchange("events-exchange", "fanout", { durable: true });

    // Enviar mensajes persistentes
    for (let i = 0; i < 5; i++) {
      const msg = { source: "producer-node", message: `event ${i}` };
      channel.publish("events-exchange", "", Buffer.from(JSON.stringify(msg)), {
        persistent: true,
      });
      console.log("Sent:", msg);
    }

    setTimeout(() => {
      connection.close();
      process.exit(0);
    }, 500);
  } catch (err) {
    console.error("Error:", err);
  }
})();
