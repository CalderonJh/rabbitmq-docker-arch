const amqp = require('amqplib');
const queue = 'events';

(async () => {
  const conn = await amqp.connect('amqp://admin:admin@rabbitmq:5672');
  const ch = await conn.createChannel();
  await ch.assertQueue(queue, {durable: true});
  console.log(' [*] Waiting for messages...');
  ch.consume(queue, msg => {
    const content = msg.content.toString();
    console.log('Received:', content);
    ch.ack(msg);
  });
})();
