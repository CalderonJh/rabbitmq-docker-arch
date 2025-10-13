const amqp = require('amqplib');
const queue = 'events';

(async () => {
  const conn = await amqp.connect('amqp://admin:admin@rabbitmq:5672');
  const ch = await conn.createChannel();
  await ch.assertQueue(queue);
  for (let i = 0; i < 5; i++) {
    const msg = JSON.stringify({ source: 'producer-node', message: `event ${i}` });
    ch.sendToQueue(queue, Buffer.from(msg));
    console.log('Sent:', msg);
  }
  await ch.close();
  await conn.close();
})();
