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
else:
    print("Could not connect to RabbitMQ.")
    exit(1)

channel = connection.channel()
channel.exchange_declare(exchange="events-exchange", exchange_type="fanout", durable=True)
channel.queue_declare(queue="events-python", durable=True)
channel.queue_bind(exchange="events-exchange", queue="events-python")
channel.basic_qos(prefetch_count=1)

def callback(ch, method, properties, body):
    print(f"Received raw: {body}")
    try:
        data = json.loads(body)
        print(f"Received from {data['source']}: {data['message']}")
    except Exception as e:
        print("Error decoding:", e)
    ch.basic_ack(delivery_tag=method.delivery_tag)

channel.basic_consume(queue="events-python", on_message_callback=callback, auto_ack=False)
print(" [*] Waiting for messages. To exit press CTRL+C")
channel.start_consuming()
