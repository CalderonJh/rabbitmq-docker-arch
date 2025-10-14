import pika, json, os, time

RABBIT_HOST = os.getenv("RABBIT_HOST", "rabbitmq")
credentials = pika.PlainCredentials("admin", "admin")

for _ in range(10):
    try:
        connection = pika.BlockingConnection(pika.ConnectionParameters(RABBIT_HOST, credentials=credentials))
        break
    except:
        print(f"Waiting for RabbitMQ at {RABBIT_HOST}:5672...")
        time.sleep(3)

channel = connection.channel()
channel.exchange_declare(exchange="events-topic-exchange", exchange_type="topic", durable=True)
channel.queue_declare(queue="events-python", durable=True)
channel.queue_bind(exchange="events-topic-exchange", queue="events-python", routing_key="events.python")

def callback(ch, method, properties, body):
    data = json.loads(body)
    print(f"[Python consumer] Received from {data['source']}: {data['message']}")
    ch.basic_ack(delivery_tag=method.delivery_tag)

channel.basic_consume(queue="events-python", on_message_callback=callback, auto_ack=False)
print(" [*] Waiting for Python events...")
channel.start_consuming()